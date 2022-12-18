package courseModel

import "managerstudent/modules/user/userModel"

type Course struct {
	SubjectTitle      string           `json:"subject_title"`
	ClassName         string           `json:"className"`
	Total             int              `json:"total"`
	MainTeacher       userModel.User   `json:"mainTeacher"`
	TeachingAssistant []userModel.User `json:"teachingAssistant"`
	TotalMale         int              `json:"totalMale"`
}
