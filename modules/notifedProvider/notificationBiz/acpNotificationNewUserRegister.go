package notificationBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notificationModel"
)

type AcpNotificationUserRegisterStore interface {
	SolveNotification(ctx context.Context, conditions interface{}, value interface{}) error
}

type acpNotificationUserRegisterBiz struct {
	store  AcpNotificationUserRegisterStore
	pubsub pubsub.Pubsub
}

func NewAcpNotificationUserRegisterBiz(store AcpNotificationUserRegisterStore, pubsub pubsub.Pubsub) *acpNotificationUserRegisterBiz {
	return &acpNotificationUserRegisterBiz{store: store, pubsub: pubsub}
}

func (biz *acpNotificationUserRegisterBiz) AcpNotifyUserRegister(ctx context.Context, data *notificationModel.Notification, status int) error {
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

	biz.pubsub.Publish(ctx, "AcpUser", pubsub.NewMessage(data))

	return nil
}
