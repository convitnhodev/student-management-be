package notifyBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notifyModel"
)

type GetNotifyStore interface {
	GetNotify(ctx context.Context, conditions interface{}) (*notifyModel.Notify, error)
}

type getNotifyBiz struct {
	store  GetNotifyStore
	pubsub pubsub.Pubsub
}

func NewGetNotifyBiz(store GetNotifyStore, pubsub pubsub.Pubsub) *getNotifyBiz {
	return &getNotifyBiz{store: store, pubsub: pubsub}
}

func (biz *getNotifyBiz) GetNotify(ctx context.Context, filter interface{}) (*notifyModel.Notify, error) {

	data, err := biz.store.GetNotify(ctx, bson.M{"id": filter})
	if err != nil {
		return nil, err
	}

	return data, err
}
