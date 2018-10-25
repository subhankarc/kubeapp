package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/smjn/kubeapp/backend/errors"
	"github.com/smjn/kubeapp/backend/models"
	"github.com/smjn/kubeapp/backend/util"
)

type UserCreator interface {
	CreateUser(models.UserModel) (*models.GeneralId, error)
}

type UserPostHandler struct {
	UDao UserCreator
}

func (u UserPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("UserPostHandler: create user")
	defer func() {
		if r := recover(); r != nil {
			log.Println("panicked")
		}
	}()

	info := models.UserModel{}
	err := json.NewDecoder(r.Body).Decode(&info)
	errors.ErrWriterPanic(w, http.StatusBadRequest, err, err, "UserPostHandler: could not create decode req body")

	id, err := u.UDao.CreateUser(info)
	errors.ErrWriterPanic(w, http.StatusInternalServerError, err, err, "UserPostHandler: could not create user")
	util.StructWriter(w, id)
}
