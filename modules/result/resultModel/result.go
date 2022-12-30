package resultModel

const NameCollection = "Results"

type Result struct {
	StudentId string    `json:"student_id" bson:"student_id,omitempty"`
	ClassId   string    `json:"class_id" bson:"class_id,omitempty"`
	CourseId  string    `json:"course_id" bson:"course_id"`
	Exam15    []float64 `json:"exam15" bson:"exam15"`
	FinalExam float64   `json:"finalExam,omitempty" bson:"finalExam"`
	Exam45    []float64 `json:"exam45" bson:"exam45"`
	QuickExam []float64 `json:"quickExam" bson:"quickExam"`
	Average   float64   `json:"average" bson:"average"`
}
