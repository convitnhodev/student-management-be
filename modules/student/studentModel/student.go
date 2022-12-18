package studentModel

import "managerstudent/modules/course/courseModel"

type Student struct {
	Id        string             `json:"id" bson:"id" :"id"`
	Gmail     string             `json:"gmail" bson:"gmail" json:"gmail,omitempty" :"gmail"`
	ClassName string             `json:"className" bson:"className"`
	Course    courseModel.Course `json:"course" bson:"course"`
	FullName  string             `json:"fullName" bson:"fullName" :"fullName"`
	Mark      Result             `json:"mark" bson:"mark"`
}
