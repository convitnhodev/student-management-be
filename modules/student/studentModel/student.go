package studentModel

const (
	StudentCollection         = "Students"
	StudentCollectionFullInfo = "FullStudents"
	Student_Class_Collection  = "student_class"
	Student_Course_Collection = "student_course"
)

type FullInfoStudent struct {
	Id        string `json:"id" bson:"id"`
	Gmail     string `json:"gmail" bson:"gmail"`
	LastName  string `json:"last_name" bson:"last_name"`
	FirstName string `json:"first_name" bson:"first_name"`
	DOB       string `json:"dob" bson:"dob"`
	Age       int    `json:"age" bson:"age"`
}

type Student struct {
	Id           string   `json:"id" bson:"id"`
	ClassId      string   `json:"class_id" bson:"class_id"`
	Results      []Result `json:"results" bson:"results"`
	AverageTotal float64  `json:"average_total" bson:"average_total"`
}
