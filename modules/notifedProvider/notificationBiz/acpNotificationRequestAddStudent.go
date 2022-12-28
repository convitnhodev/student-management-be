package notificationBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notificationModel"
)

type AcpNotificationRequestAddStudentStore interface {
	SolveNotification(ctx context.Context, conditions interface{}, value interface{}) error
}

type acpNotificationRequestAddStudentBiz struct {
	store  AcpNotificationRequestAddStudentStore
	pubsub pubsub.Pubsub
}

func NewAcpNotificationRequestAddStudentBiz(store AcpNotificationRequestAddStudentStore, pubsub pubsub.Pubsub) *acpNotificationRequestAddStudentBiz {
	return &acpNotificationRequestAddStudentBiz{store: store, pubsub: pubsub}
}

func (biz *acpNotificationRequestAddStudentBiz) AcpNotificationRequestAddStudent(ctx context.Context, data *notificationModel.Notification, status int) error {
	data.Status = status

	filter := bson.D{{"_id", data.Id}}
	update := bson.D{{"$set", bson.D{{"status", status}, {"seen", true}}}}
	err := biz.store.SolveNotification(ctx, filter, update)
	if err != nil {
		return err
	}
	if status == 0 {
		return nil
	}

	biz.pubsub.Publish(ctx, "AcpStudent", pubsub.NewMessage(data))

	return nil
}
