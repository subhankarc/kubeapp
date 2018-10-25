package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/smjn/kubeapp/backend/db"
	"github.com/smjn/kubeapp/backend/errors"
	"github.com/smjn/kubeapp/backend/models"
	"github.com/smjn/kubeapp/backend/util"
)

type UserPutHandler struct {
}

var (
	errAliasInvalid = fmt.Errorf("alias cannot be more than 10 chars and should be alphanumeric")

	reAlias, _ = regexp.Compile("^[0-9a-zA-Z]{1,10}$")
)

func (p UserPutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("UserPutHandler: new update user profile request")
	defer func() {
		if r := recover(); r != nil {
			log.Println("panicked")
		}
	}()

	pathVar := mux.Vars(r)
	if _, ok := pathVar["inumber"]; !ok {
		errors.ErrWriterPanic(w, http.StatusBadRequest, errors.ErrUserNotFound, errors.ErrUserNotFound, "UserPutHandler: error inumber not provided")
	}
	err := p.parseAndUpdate(r, pathVar["inumber"])
	errors.ErrWriterPanic(w, http.StatusBadRequest, err, err, "UserPutHandler: error parsing req data")
	util.OkWriter(w)
}

func (p UserPutHandler) parseAndUpdate(r *http.Request, inumber string) error {
	log.Println("UserPutHandler: parsing request")
	query := "update myuser set"
	values := []interface{}{}

	defer r.Body.Close()
	info := &models.UserModel{}
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		log.Println("UserPutHandler: ", err)
		return err
	}

	ind := 1
	if info.Firstname != "" {
		values = append(values, info.Firstname)
		query += fmt.Sprintf(" firstname=$%d,", ind)
		ind++
	}
	if info.Lastname != "" {
		values = append(values, info.Lastname)
		query += fmt.Sprintf(" lastname=$%d,", ind)
		ind++
	}
	if info.Alias != "" {
		values = append(values, info.Alias)
		query += fmt.Sprintf(" alias=$%d,", ind)
		ind++
	}

	if ind > 1 {
		query = query[0 : len(query)-1]
	} else {
		//no updates
		return nil
	}

	values = append(values, inumber)
	query += fmt.Sprintf(" where inumber=$%d", ind)

	log.Println(query, values)
	_, err := db.DB.Exec(query, values...)
	return err
}
