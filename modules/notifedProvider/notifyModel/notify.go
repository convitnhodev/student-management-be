package notifyModel

import "time"

type Notify struct {
	Id      string    `json:"id" bson:"id"`
	Content string    `json:"content" json:"content"`
	Agent   string    `json:"agent" json:"agent"`
	Passive string    `json:"passive" json:"passive"`
	Time    time.Time `json:"time" json:"time"`
	Seen    bool      `json:"seen" bson:"seen"`
}
