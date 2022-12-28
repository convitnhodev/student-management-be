package courseModel

const NameCollection = "Courses"

type Course struct {
	Id           string `json:"course_id,omitempty" bson:"course_id"`
	SubjectTitle string `json:"subject_title" bson:"subjectTitle"`
	ClassId      string `json:"class_id" bson:"class_id"`
	Total        int    `json:"total" bson:"total"`
	MainTeacher  string `json:"main_teacher" bson:"main_teacher"`
}
