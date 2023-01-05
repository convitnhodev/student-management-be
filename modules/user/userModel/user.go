package userModel

import (
	_const "managerstudent/common/const"
	"time"
)

const NameCollection = "Users"

type User struct {
	UserName string      `json:"username,omitempty" bson:"username,omitempty"`
	Password string      `json:"password,omitempty" bson:"password,omitempty"`
	FullName string      `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Class    string      `json:"class,omitempty" bson:"class,omitempty"`
	Role     _const.Role `json:"role,omitempty" bson:"role,omitempty"`
	Salt     string      `json:"-" bson:"salt,omitempty"`
	Token    string      `json:"-" bson:"token,omitempty"`
	Phone    string      `json:"phone" bson:"phone,omitempty"`
	Gmail    string      `json:"gmail" bson:"gmail,omitempty"`
	Acp      bool        `json:"acp" bson:"acp,omitempty"`
	Sex      bool        `json:"sex" bson:"sex,omitempty"`
	Address  string      `json:"address" bson:"address,omitempty"`
	Dob      time.Time   `json:"dob" bson:"dob,omitempty"`
}

func (user *User) GetUserName() string {
	return user.UserName
}

type UpdatePassWord struct {
	UserName    string `json:"username" bson:"username,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	NewPassword string `json:"new_password"`
}
