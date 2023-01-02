package resultModel

type Filter struct {
	IdStudent *string `json:"student_id" bson:"student_id" form:"student_id"`
	IdClass   *string `json:"class_id" bson:"class_id" form:"class_id"`
	IdCourse  *string `json:"course_id" bson:"course_id" form:"course_id"`
}
