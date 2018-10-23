package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smjn/ipl18/backend/errors"
	"github.com/smjn/ipl18/backend/util"
)

type UserDeleter interface {
	DeleteUserById(inumber string) error
}

type UserDeleteHandler struct {
	UDao UserDeleter
}

func (u UserDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("UserDeleteHandler: request to delete profile")
	defer func() {
		if r := recover(); r != nil {
			log.Println("panicked")
		}
	}()

	pathVars := mux.Vars(r)
	if _, ok := pathVars["inumber"]; !ok {
		errors.ErrWriterPanic(w, http.StatusBadRequest, errors.ErrUserNotFound, errors.ErrUserNotFound, "UserDeleteHandler: user id not set in path")
	}

	if err := u.UDao.DeleteUserById(pathVars["inumber"]); err != nil {
		errors.ErrWriterPanic(w, http.StatusInternalServerError, errors.ErrDBIssue, errors.ErrDBIssue, "UserDeleteHandler: could not delete user")
	}
	util.OkWriter(w)
	return
}
