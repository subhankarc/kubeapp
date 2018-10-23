package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smjn/ipl18/backend/errors"
	"github.com/smjn/ipl18/backend/models"
	"github.com/smjn/ipl18/backend/util"
)

type UserGetter interface {
	GetUsers() (*models.UsersModel, error)
	GetUserById(inumber string) (*models.UserModel, error)
}

type UserGetHandler struct {
	UDao UserGetter
}

const (
	qFetchUserDetails = "select firstname, lastname, alias, inumber ipluser where inumber=$1"
)

var (
	errUserNotFound = fmt.Errorf("user not found")
)

func (u UserGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("UserGetHandler: request to get user(s)")
	defer func() {
		if r := recover(); r != nil {
			log.Println("panicked")
		}
	}()

	pathVar := mux.Vars(r)
	if _, ok := pathVar["inumber"]; !ok {
		//fetch all users here
		if info, err := u.UDao.GetUsers(); err != nil {
			errors.ErrWriterPanic(w, http.StatusInternalServerError, err, errors.ErrDBIssue, "UserGetHandler: could not query db")
		} else {
			log.Println("UserGetHandler:", info)
			util.StructWriter(w, info)
		}
		return
	}

	//fetch single user here
	if info, err := u.UDao.GetUserById(pathVar["inumber"]); err != nil {
		errors.ErrWriterPanic(w, http.StatusNotFound, err, errUserNotFound, "UserGetHandler: user not found in db")
	} else {
		log.Println("UserGetHandler:", info)
		util.StructWriter(w, &info)
	}
}
