package subjectModel


const NameCollection = "Subjects"

type Subject struct {
	ClassID string `json:"class_id" bson:"class_id"`
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
}
