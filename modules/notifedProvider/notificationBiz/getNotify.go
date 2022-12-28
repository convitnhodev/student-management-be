package notificationBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notificationModel"
)

type GetNotifyStore interface {
	GetNotification(ctx context.Context, conditions interface{}) (*notificationModel.Notification, error)
}

type getNotifyBiz struct {
	store  GetNotifyStore
	pubsub pubsub.Pubsub
}

func NewGetNotifyBiz(store GetNotifyStore, pubsub pubsub.Pubsub) *getNotifyBiz {
	return &getNotifyBiz{store: store, pubsub: pubsub}
}

func (biz *getNotifyBiz) GetNotification(ctx context.Context, filter interface{}) (*notificationModel.Notification, error) {

	data, err := biz.store.GetNotification(ctx, bson.M{"_id": filter})
	if err != nil {
		return nil, err
	}

	return data, err
}
