package courseModel

const NameCollection = "Courses"

type Course struct {
	Id            string   `json:"course_id,omitempty" bson:"course_id"`
	SubjectTitle  string   `json:"subject_title" bson:"subjectTitle"`
	ClassId       string   `json:"class_id" bson:"class_id"`
	Total         int      `json:"total" bson:"total"`
	MainTeacher   string   `json:"main_teacher" bson:"main_teacher"`
	ListStudentId []string `json:"list_student_id" bson:"list_student_id"`
	MaxAge        int      `json:"max_age" bson:"max_age"`
	MaxTotal      int      `json:"max_total" bson:"max_total"`
	MinMark       float64  `json:"min_mark" bson:"min_mark"`
}


