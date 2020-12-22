package model

import (
	"encoding/json"
	"errors"
	"go-gin-web/pkg/errno"
	"log"
)

type User struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (user *User)SelectUserByName(name string)error {
	stmt,err := DB.Prepare("SELECT user_name,password FROM user WHERE user_name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	defer rows.Close()
	if err != nil {
		return err
	}
	// 数据处理
	for rows.Next() {
		rows.Scan( &user.Username, &user.Password)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

// Validate the fields.
func (u *User) Validate() error {
	if u.Username =="" || u.Password ==""{
		return errors.New(errno.ErrValidation.Message)
	}
	return nil
}
func (user *User) Create() (int64,error)  {
	id,err := Insert("INSERT INTO  user(user_name,password) values (?,?)", user.Username, user.Password)
	if err != nil {
		return 1,err
	}

	return id,nil
}

func (user *User)UserToJson()string  {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}

func (user *User)JsonToUser(jsonBlob string)error  {
	err := json.Unmarshal([]byte(jsonBlob), &user)
	if err != nil {
		return err
	}
	return nil
}