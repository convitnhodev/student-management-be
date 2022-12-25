package classModel

type Class struct {
	Id              string `json:"class_id" bson:"class_id"`
	Grade           int    `json:"grade"`
	Name            string `json:"name"`
	Total           int    `json:"total"`
	TotalMale       int    `json:"total_male"`
	HomeroomTeacher string `json:"homeroomTeacher"`
}
