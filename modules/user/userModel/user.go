package userModel

import (
	_const "managerstudent/common/const"
	"time"
)

const NameCollection = "Users"

type User struct {
	UserName string      `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Password string      `json:"pass_word,omitempty" bson:"pass_word,omitempty"`
	FullName string      `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Class    string      `json:"class,omitempty" bson:"class,omitempty"`
	Role     _const.Role `json:"role,omitempty" bson:"role,omitempty"`
	School   string      `json:"school,omitempty" bson:"school,omitempty"`
	Salt     string      `json:"-" bson:"salt,omitempty"`
	Token    string      `json:"-" bson:"token,omitempty"`
	Phone    string      `json:"phone" bson:"phone,omitempty"`
	Gmail    string      `json:"gmail" bson:"gmail,omitempty"`
	Acp      bool        `json:"acp" bson:"acp,omitempty"`
	sex      bool        `json:"sex" bson:"sex,omitempty""`
	address  string      `json:"address" bson:"address,omitempty"`
	dob      time.Time   `json:"dob" bson:"dob,omitempty"`
}

func (user *User) GetUserName() string {
	return user.UserName
}

type UpdatePassWord struct {
	UserName    string `json:"user_name" bson:"user_name,omitempty"`
	Password    string `json:"pass_word" bson:"pass_word,omitempty"`
	NewPassword string `json:"new_pass_word"`
}
