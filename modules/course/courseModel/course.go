package courseModel

import "managerstudent/modules/user/userModel"

type Course struct {
	Id                string           `json:"course_id,omitempty " bson:"course_id"`
	SubjectTitle      string           `json:"subject_title" bson:"subjectTitle"`
	ClassName         string           `json:"className" bson:"className"`
	Total             int              `json:"total" bson:"total"`
	MainTeacher       userModel.User   `json:"mainTeacher" bson:"mainTeacher"`
	TeachingAssistant []userModel.User `json:"teachingAssistant" bson:"teachingAssistant"`
	TotalMale         int              `json:"totalMale" bson:"totalMale,omitempty"`
}
