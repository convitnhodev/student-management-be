package notifyModel

import "time"

type Notify struct {
	Id      int       `json:"id" bson:"id"`
	Content string    `json:"content" bson:"content"`
	Agent   string    `json:"agent" bson:"agent"`
	Passive string    `json:"passive" bson:"passive"`
	Time    time.Time `json:"time" bson:"time"`
	Seen    bool      `json:"seen" bson:"seen"`
	Status  int       `bson:"status" bson:"status"`
}
