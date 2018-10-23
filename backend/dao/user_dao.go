package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/smjn/ipl18/backend/db"
	"github.com/smjn/ipl18/backend/errors"
	"github.com/smjn/ipl18/backend/models"
)

type UserDAO struct{}

const (
	qSelectAllUsers = "select firstname,lastname,alias,inumber from myuser"
	qSelectUserById = "select firstname,lastname,alias,inumber from myuser where inumber=$1"
	qInsertUser     = "insert into myuser(firstname, lastname, alias, inumber) values($1, $2, $3, $4)"
	qDeleteUserById = "delete from myuser where inumber=$1"
)

var errLeaderNotFound = fmt.Errorf("leader not found in db")

func (u UserDAO) CreateUser(user models.UserModel) (*models.GeneralId, error) {
	res, err := db.DB.Exec(qInsertUser,
		user.Firstname,
		user.Lastname,
		user.Alias,
		user.INumber)

	if err != nil {
		log.Println("UserDAO: InsertUser error inserting new user", err)
		return nil, &errors.DaoError{http.StatusInternalServerError, err, errors.ErrDBIssue}
	}

	if num, err := res.RowsAffected(); err != nil || num != 1 {
		log.Println("UserDAO: db not updated inserting new user", err, num)
		return nil, &errors.DaoError{http.StatusInternalServerError, errors.ErrDBIssue, errors.ErrDBIssue}
	}

	return &models.GeneralId{user.INumber}, nil
}

func (u UserDAO) GetUsers() (*models.UsersModel, error) {
	rows, err := db.DB.Query(qSelectAllUsers)
	if err != nil {
		log.Println("UserDAO: GetUsers error getting users", err)
		return nil, &errors.DaoError{http.StatusInternalServerError, err, errors.ErrDBIssue}
	}

	users := []*models.UserModel{}
	defer rows.Close()
	for rows.Next() {
		user := models.UserModel{}
		err := rows.Scan(&user.Firstname, &user.Lastname, &user.Alias, &user.INumber)
		if err == sql.ErrNoRows {
			log.Println("UserDAO: GetUsers could not find users", err)
			return nil, &errors.DaoError{http.StatusNotFound, err, errLeaderNotFound}
		} else if err != nil {
			log.Println("UserDAO: GetUsers error scanning users", err)
			return nil, &errors.DaoError{http.StatusInternalServerError, err, errors.ErrDBIssue}
		}
		users = append(users, &user)
	}

	return &models.UsersModel{users}, nil
}

func (u UserDAO) GetUserById(inumber string) (*models.UserModel, error) {
	info := models.UserModel{}
	if err := db.DB.QueryRow(qSelectUserById, inumber).Scan(&info.Firstname, &info.Lastname, &info.Alias, &info.INumber); err == sql.ErrNoRows {
		log.Println("UserDAO: GetUserById could not find user", err)
		return nil, &errors.DaoError{http.StatusBadRequest, errors.ErrUserNotFound, errors.ErrUserNotFound}
	} else if err != nil {
		return nil, &errors.DaoError{http.StatusInternalServerError, errors.ErrUserNotFound, errors.ErrUserNotFound}
	}
	return &info, nil
}

func (u UserDAO) DeleteUserById(inumber string) error {
	if res, err := db.DB.Exec(qDeleteUserById, inumber); err != nil {
		log.Println("UserDAO: DeleteUserById error deleting user", err)
		return &errors.DaoError{http.StatusInternalServerError, err, errors.ErrDBIssue}
	} else {
		if i, err := res.RowsAffected(); err != nil || i != 1 {
			log.Println("UserDAO: DeleteUserById affected rows don't add up", err, i)
			return &errors.DaoError{http.StatusInternalServerError, errors.ErrDBIssue, errors.ErrDBIssue}
		}
	}
	return nil
}
