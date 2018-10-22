package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/smjn/ipl18/backend/db"
	"github.com/smjn/ipl18/backend/errors"
	"github.com/smjn/ipl18/backend/util"
)

type UserPutHandler struct {
}

const (
	maxMemory = 1024 * 1024 * 2
)

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
	if _,ok:=pathVar["inumber"]; !ok{
	    errors.ErrWriterPanic(w, http.StatusBadRequest, http.StatusBadRequest, http.StatusBadRequest, "UserPutHandler: error inumber not provided")
	}
	err = p.parseAndUpdate(r, pathVar["inumber"])
	errors.ErrWriterPanic(w, http.StatusBadRequest, err, err, "UserPutHandler: error parsing form data")
	util.OkWriter(w)
}

func (p UserPutHandler) parseAndUpdate(r *http.Request, inumber string) error {
	log.Println("UserPutHandler: parsing request")
	query := "update ipluser"
	values := []interface{}{inumber}
	i := 2

	defer r.Body.Close()
	info:=&models.UserModel{}
	if err:=json.NewDecoder(r.Body).Decode(&info);err!=nil{
		log.Println("ProfileDAO: ",err)
		return err
	}

	ind:=1
	if(info.FirstName!=""){
		values=values.append(info.Firstname)
		query+=fmt.Sprintf(" set firstname=$%d,",ind)
		ind++
	}
	if(info.Lastname!=""){
		values=values.append(info.Lastname)
		query+=fmt.Sprintf(" set lastname=$%d,",ind)
		ind++
	}
	if(info.Alias!=""){
		values=values.append(info.Alias)
		query+=fmt.Sprintf(" set alias=$%d,",ind)
		ind++
	}
	if query[len(query)-1]==","{
		query=query[0:len(query)-1]
	}else{
		//no updates
		return nil
	}
	
	query += fmt.Sprintf(" where inumber=$%d", ind)

	log.Println(query, values)
	_, err := db.DB.Exec(query, values...)
	return err
}
