package classModel

type Class struct {
	Id                      string `json:"class_id" bson:"class_id"`
	Grade                   int    `json:"grade"`
	Name                    string `json:"name"`
	Total                   int    `json:"total"`
	TotalMale               int    `json:"totalMale"`
	TotalEthnicMinority     int    `json:"totalEthnicMinority"`
	TotalMaleEthnicMinority int    `json:"totalMaleEthnicMinority"`
	HomeroomTeacher         string `json:"homeroomTeacher"`
}
