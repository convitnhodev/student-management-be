package userModel

import _const "managerstudent/common/const"

type UserLogin struct {
	UserName string      `json:"user_name" bson:"user_name"`
	Password string      `json:"pass_word" bson:"pass_word"`
	Role     _const.Role `json:"role,omitempty" bson:"role"`
}
