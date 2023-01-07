package studentModel

import "time"

const (
	NameCollection = "Students"
)

type Student struct {
	Id       string    `json:"id" bson:"id"`
	FullName string    `json:"fullname" bson:"fullname"`
	ClassId  string    `json:"class_id" bson:"class_id"`
	Email    string    `json:"email" bson:"email"`
	DOB      time.Time `json:"dob" bson:"dob"`
	Sex      bool      `json:"sex" bson:"sex"`
	Phone    string    `json:"phone" bson:"phone"`
	Address  string    `json:"address" bson:"address"`
}
