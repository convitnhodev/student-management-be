package courseModel

type Filter struct {
	SubjectTitle string `json:"subject_title,omitempty" bson:"subject_title,omitempty"`
	MainTeacher  string `json:"main_teacher,omitempty" bson:"main_teacher,omitempty"`
}
