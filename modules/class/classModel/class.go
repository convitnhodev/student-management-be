package classModel

type Class struct {
	Grade                   int    `json:"grade"`
	Name                    string `json:"name"`
	Total                   int    `json:"total"`
	TotalMale               int    `json:"totalMale"`
	TotalEthnicMinority     int    `json:"totalEthnicMinority"`
	TotalMaleEthnicMinority int    `json:"totalMaleEthnicMinority"`
	homeroomTeacher         string `json:"homeroomTeacher"`
}
