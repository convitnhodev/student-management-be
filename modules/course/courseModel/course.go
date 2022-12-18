package courseModel

import "managerstudent/modules/user/userModel"

type Course struct {
	Id                string           `json:"id,omitempty"`
	SubjectTitle      string           `json:"subject_title" json:"subjectTitle,omitempty"`
	ClassName         string           `json:"className" json:"className,omitempty"`
	Total             int              `json:"total" json:"total,omitempty"`
	MainTeacher       userModel.User   `json:"mainTeacher" json:"mainTeacher"`
	TeachingAssistant []userModel.User `json:"teachingAssistant" json:"teachingAssistant,omitempty"`
	TotalMale         int              `json:"totalMale" json:"totalMale,omitempty"`
}
