package resultModel

const NameCollection = "Results"

type Result struct {
	StudentId string    `json:"student_id" bson:"student_id,omitempty"`
	ClassId   string    `json:"class_id" bson:"class_id,omitempty"`
	CourseId  string    `json:"course_id" bson:"course_id"`
	Exam15    []float64 `json:"exam15" bson:"exam15"`
	FinalExam *float64  `json:"final_exam" bson:"final_exam"`
	Exam45    []float64 `json:"exam45,omitempty" bson:"exam45"`
	Average   float64   `json:"average" bson:"average"`
}

func (r *Result) CalculateAverage() {
	var sum float64
	for _, v := range r.Exam15 {
		sum += v
	}
	for _, v := range r.Exam45 {
		sum += (2 * v)
	}
	if r.FinalExam != nil {
		sum += (3 * *r.FinalExam)
	}

	coefficient := (float64(len(r.Exam15)) + 2*float64(len(r.Exam45)) + 3)
	if r.FinalExam == nil {
		coefficient -= 3
	}
	r.Average = sum / coefficient
}
