package resultModel

type Filter struct {
	IdStudent string `json:"id_student" bson:"id_student" form:"id_student"`
	IdClass   string `json:"id_class" bson:"id_class" form:"id_class"`
	IdCourse  string `json:"id_course" bson:"id_course" form:"id_class"`
}
