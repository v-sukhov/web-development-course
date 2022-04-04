package db

import (
	"database/sql"
	"log"
	"strings"
)

type UserInfo struct {
	UserId      int
	UserLogin   string
	DisplayName string
	IsAdmin     int
	IsTeacher   int
}

func AuthenticateUser(login string, password string) (UserInfo, bool) {
	userInfo := UserInfo{}
	ok := true
	err := db.QueryRow("SELECT user_id, login, COALESCE(display_name, '') display_name, is_admin, is_teacher FROM t_user WHERE login = $1 and password_sha256 = sha256($2)",
		login, EncryptionSaltWord+password).
		Scan(&userInfo.UserId, &userInfo.UserLogin, &userInfo.DisplayName, &userInfo.IsAdmin, &userInfo.IsTeacher)
	if err != nil {
		ok = false
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	} else {
		ok = true
	}

	return userInfo, ok
}

func FindUserByLogin(login string) (UserInfo, bool) {
	userInfo := UserInfo{}
	ok := true
	err := db.QueryRow("SELECT user_id, login, COALESCE(display_name, '') display_name, is_admin, is_teacher FROM t_user WHERE LOWER(login) = $1 ",
		strings.ToLower(login)).
		Scan(&userInfo.UserId, &userInfo.UserLogin, &userInfo.DisplayName, &userInfo.IsAdmin, &userInfo.IsTeacher)

	if err != nil {
		ok = false
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	} else {
		ok = true
	}

	return userInfo, ok
}

func CreateUser(userLogin string, email string, password string, is_admin int, is_teacher int) (int, error) {

	var userId int

	err := db.QueryRow(`INSERT INTO t_user(login, email, password_sha256, is_admin, is_teacher) 
						VALUES( $1, $2, sha256($3), $4, $5) 
						RETURNING user_id`, userLogin, email, EncryptionSaltWord+password, is_admin, is_teacher).Scan(&userId)
	if err != nil {
		log.Fatal(err)
	}

	return userId, err
}

func UpdatePassword(userLogin string, password string) {
	db.Exec(`UPDATE t_user SET password_sha256 = sha256($2) WHERE login = $1`, userLogin, EncryptionSaltWord+password)
}
