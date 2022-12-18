package userModel

import (
	"errors"
	_const "managerstudent/common/const"
	"strings"
)

type User struct {
	UserName string      `json:"user_name,omitempty" bson:"user_name"`
	Password string      `json:"password,omitempty" bson:"password"`
	FullName string      `json:"full_name,omitempty" bson:"full_name"`
	Class    string      `json:"class,omitempty" bson:"class"`
	Role     _const.Role `json:"role,omitempty" bson:"role"`
	School   string      `json:"school,omitempty" bson:"school"`
	Salt     string      `json:"-" bson:"salt"`
	Token    string      `json:"-" bson:"token"`
	Phone    string      `json:"phone" bson:"phone"`
	Gmail    string      `json:"gmail" bson:"gmail"`
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
		return errors.New("full_name name can not be blank")
	}

	return nil
}

func (user *User) GetUserName() string {
	return user.UserName
}
