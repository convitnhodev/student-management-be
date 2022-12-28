package resultModel

const NameCollection = "Results"

type Result struct {
	IdStudent string    `json:"id_student" bson:"id_student,omitempty"`
	IdClass   string    `json:"id_class" bson:"id_class,omitempty"`
	IdCourse  string    `json:"id_course" bson:"id_course,omitempty"`
	Exam15    []float64 `json:"exam15" bson:"exam15,omitempty"`
	FinalExam float64   `json:"finalExam,omitempty" bson:"finalExam,omitempty"`
	Exam45    float64   `json:"exam45" bson:"exam45,omitempty"`
	QuickExam float64   `json:"quickExam" bson:"quickExam,omitempty"`
	Average   float64   `json:"average" bson:"average,omitempty"`
}

type ResultUpdate struct {
	Exam15    []float64 `json:"exam15" bson:"exam15"`
	FinalExam float64   `json:"finalExam,omitempty" bson:"finalExam"`
	Exam45    float64   `json:"exam45" bson:"exam45"`
	QuickExam float64   `json:"quickExam" bson:"quickExam"`
	Average   float64   `json:"average" bson:"average"`
}
