package userModel

import (
	_const "managerstudent/common/const"
	"time"
)

const NameCollection = "Users"

type User struct {
	UserName        string      `json:"username" bson:"username"`
	Password        string      `json:"password" bson:"password"`
	FullName        string      `json:"fullname" bson:"fullname"`
	Class           string      `json:"class" bson:"class"`
	TeachingClasses []string    `json:"teaching_classes" bson:"teaching_classes"`
	Role            _const.Role `json:"role" bson:"role"`
	Salt            string      `json:"-" bson:"salt"`
	Token           string      `json:"-" bson:"token"`
	Phone           string      `json:"phone" bson:"phone"`
	Gmail           string      `json:"gmail" bson:"gmail"`
	Acp             bool        `json:"acp" bson:"acp"`
	Sex             bool        `json:"sex" bson:"sex"`
	Address         string      `json:"address" bson:"address"`
	Dob             time.Time   `json:"dob" bson:"dob"`
}

func (user *User) GetUserName() string {
	return user.UserName
}

func (user *User) GetRole() string {
	return user.Role.String()
}

type UpdatePassWord struct {
	UserName    string `json:"username" bson:"username,omitempty"`
	Password    string `json:"password" bson:"password,omitempty"`
	NewPassword string `json:"new_password"`
}

type UpdateHomeroom struct {
	UserName string `json:"username" bson:"username,omitempty"`
	Class    string `json:"class" bson:"class,omitempty"`
}

type UpdateTeaching struct {
	UserName        string   `json:"username" bson:"username,omitempty"`
	TeachingClasses []string `json:"teaching_classes" bson:"teaching_classes,omitempty"`
}
