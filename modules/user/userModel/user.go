package userModel

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_const "managerstudent/common/const"
	"strings"
	"time"
)

const NameCollection = "Users"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName string             `json:"user_name,omitempty" bson:"user_name"`
	Password string             `json:"pass_word,omitempty" bson:"password"`
	FullName string             `json:"full_name,omitempty" bson:"full_name"`
	Class    string             `json:"class,omitempty" bson:"class"`
	Role     _const.Role        `json:"role,omitempty" bson:"role"`
	School   string             `json:"school,omitempty" bson:"school"`
	Salt     string             `json:"-" bson:"salt"`
	Token    string             `json:"-" bson:"token"`
	Phone    string             `json:"phone" bson:"phone"`
	Gmail    string             `json:"gmail" bson:"gmail"`
	Acp      bool               `json:"acp" bson:"acp"`
	sex      bool               `json:"sex" bson:"sex""`
	address  string             `json:"address" bson:"address"`
	dob      time.Time          `json:"dob" bson:"dob"`
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
