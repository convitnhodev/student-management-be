package classModel

const NameCollection = "Classes"

type Class struct {
	Id              string   `json:"class_id" bson:"class_id"`
	Grade           int      `json:"grade" bson:"grade"`
	Total           int      `json:"total" bson:"total"`
	TotalMale       int      `json:"total_male" bson:"total_male"`
	HomeroomTeacher string   `json:"homeroom_teacher" bson:"homeroom_teacher"`
	ListStudentId   []string `json:"list_student_id" bson:"list_student_id"`
}
