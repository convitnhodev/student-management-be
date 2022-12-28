package notifyModel

import "time"

type Notification struct {
	Id       int       `json:"id" bson:"id"`
	Content  string    `json:"content" bson:"content"`
	Agent    string    `json:"agent" bson:"agent"`
	Passive  string    `json:"passive" bson:"passive"`
	Location string    `json:"location" bson:"location"`
	Time     time.Time `json:"time" bson:"time"`
	Seen     bool      `json:"seen" bson:"seen"`
	Status   int       `bson:"status" bson:"status"`
}