package classModel

type Filter struct {
	Grade           int    `json:"grade,omitempty" bson:"grade,omitempty"`
	HomeroomTeacher string `json:"homeroom_teacher" bson:"homeroom_teacher"`
}
