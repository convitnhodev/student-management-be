package studentModel

type StudentAndClass struct {
	StudentId string `json:"student_id" bson:"student_id"`
	ClassId   string `json:"class_id" bson:"class_id"`
}

type StudentAndCourse struct {
	StudentId string   `json:"student_id" bson:"student_id"`
	Courses   []string `json:"list_course_id" bson:"list_course_id"`
}
