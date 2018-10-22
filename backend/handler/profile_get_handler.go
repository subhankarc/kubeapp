package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/smjn/ipl18/backend/db"
	"github.com/smjn/ipl18/backend/errors"
	"github.com/smjn/ipl18/backend/models"
	"github.com/smjn/ipl18/backend/util"
)

// UserGetHandler .
type UserGetHandler struct {
}

const (
	qFetchUserDetails = "select firstname, lastname, alias, inumber ipluser where inumber=$1"
)

var (
	errUserNotFound = fmt.Errorf("user not found")
)

func (p UserGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("UserGetHandler: request to get profile")
	defer func() {
		if r := recover(); r != nil {
			log.Println("panicked")
		}
	}()

	pathVar := mux.Vars(r)
	if _,ok:=pathVar["inumber"];!ok {
		//fetch all users here
	}

	//fetch single user here
	info := models.UserModel{}

	if err := db.DB.QueryRow(qFetchUserDetails, inumber).Scan(&info.Firstname, &info.Lastname, &info.Alias, &info.INumber); err == sql.ErrNoRows {
		errors.ErrWriterPanic(w, http.StatusForbidden, err, errUserNotFound, "UserGetHandler: user not found in db")
	}
	errors.ErrWriterPanic(w, http.StatusInternalServerError, err, errors.ErrDBIssue, "UserGetHandler: could not query db")

	log.Println("UserGetHandler:", info)
	util.StructWriter(w, &info)
}
