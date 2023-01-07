package subjectModel


const NameCollection = "Subjects"

type Subject struct {
	ID    string `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
}
