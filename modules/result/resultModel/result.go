package resultModel

const NameCollection = "Results"

type Result struct {
	StudentId string  `json:"student_id" bson:"student_id,omitempty"`
	ClassId   string  `json:"class_id" bson:"class_id,omitempty"`
	SubjectId string  `json:"subject_id" bson:"subject_id,omitempty"`
	Exam15    float64 `json:"exam15" bson:"exam15"`
	Exam45    float64 `json:"exam45,omitempty" bson:"exam45"`
	FinalExam float64 `json:"final_exam" bson:"final_exam"`
	Average   float64 `json:"average" bson:"average"`
}

func (r *Result) CalculateAverage() {
	r.Average = (r.Exam15 + r.Exam45*2 + r.FinalExam*3) / 6
}
