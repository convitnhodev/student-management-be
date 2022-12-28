package notificationModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const NameCollection = "Notifications"

type Notification struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content  string             `json:"content" bson:"content"`
	Agent    string             `json:"agent" bson:"agent"`
	Passive  string             `json:"passive" bson:"passive"`
	Location string             `json:"location" bson:"location"`
	Time     time.Time          `json:"time" bson:"time"`
	Seen     bool               `json:"seen" bson:"seen"`
	Status   int                `bson:"status" bson:"status"`
}
