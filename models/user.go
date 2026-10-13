package models

import (
	"errors"

	"github.com/ahmed0naser/event-manager-go/db"
	"github.com/ahmed0naser/event-manager-go/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save()error {
	query := `INSERT INTO users(email,password) VALUES (?,?)`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
return err
	}
	defer stmt.Close()
	hashedPass,err:=utils.HashPassword(u.Password)
	if err!=nil{
return err
	}
res,err:=stmt.Exec(u.Email,hashedPass)
	if err!=nil{
return err
	}
	userId,err:=res.LastInsertId()
	u.ID=userId

	return err
}

func (user *User) ValidateCredentials() error{
	query:=`SELECT id,password FROM users WHERE email=?`
	row:=db.DB.QueryRow(query,user.Email)
	var password string
	err:=row.Scan(&user.ID,&password)
	if err!=nil{
		return err
	}
	passwordIsValid:=utils.ComparePassword(password,user.Password)
	if !passwordIsValid {
		return errors.New("Invalid Credentials")
	}
	return nil
}