package userModel

import _const "managerstudent/common/const"

type UserLogin struct {
	UserName string      `json:"user_name,omitempty" bson:"user_name"`
	Password string      `json:"password,omitempty" bson:"password"`
	Role     _const.Role `json:"role,omitempty" bson:"role"`
}
