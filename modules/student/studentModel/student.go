package studentModel

const (
	StudentCollection         = "Students"
	Student_Class_Collection  = "student_class"
	Student_Course_Collection = "student_course"
)

type Student struct {
	Id       string   `json:"id" bson:"id"`
	Gmail    string   `json:"gmail" bson:"gmail" json:"gmail,omitempty"`
	ClassId  string   `json:"class_id" bson:"class_id"`
	CourseId []string `json:"course_id" bson:"course_id"`
	FullName string   `json:"full_name" bson:"full_name"`
	Acp      bool     `json:"acp" bson:"acp"`
}
