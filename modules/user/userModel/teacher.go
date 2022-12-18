package userModel

import (
	"errors"
	_const "managerstudent/common/const"
	"strings"
)

type User struct {
	UserName string      `json:"user_name" bson:"user_name,omitempty"`
	Password string      `json:"password" bson:"password"`
	FullName string      `json:"full_name" bson:"full_name"`
	Class    string      `bson:"class,omitempty"`
	Role     _const.Role `json:"role" bson:"role"`
}

func (user *User) Validate() error {

	//check validate of email
	user.UserName = strings.TrimSpace(user.UserName)
	user.FullName = strings.TrimSpace(user.FullName)
	user.Class = strings.TrimSpace(user.Class)
	user.Password = strings.TrimSpace(user.Password)

	if user.UserName == "" {
		return errors.New("username name can not be blank")
	}

	if user.Password == "" {
		return errors.New("password name can not be blank")
	}

	if user.FullName == "" {
		return errors.New("fullname name can not be blank")
	}

	return nil
}

func (user *User) GetUserName() string {
	return user.UserName
}
