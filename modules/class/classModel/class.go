package classModel

type Class struct {
	Id              string `json:"class_id" bson:"class_id"`
	Grade           int    `json:"grade" bson:"grade"`
	Name            string `json:"name" bson:"name"`
	Total           int    `json:"total" bson:"total"`
	TotalMale       int    `json:"total_male" bson:"total_male"`
	HomeroomTeacher string `json:"homeroom_teacher" bson:"homeroom_teacher"`
}
