package notifyBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notifyModel"
)

type AcpNotifyStore interface {
	SolveNotify(ctx context.Context, conditions interface{}, value interface{}) error
}

type acpNotifyBiz struct {
	store  AcpNotifyStore
	pubsub pubsub.Pubsub
}

func NewAcpNotifyBiz(store AcpNotifyStore, pubsub pubsub.Pubsub) *acpNotifyBiz {
	return &acpNotifyBiz{store: store, pubsub: pubsub}
}

func (biz *acpNotifyBiz) AcpNotify(ctx context.Context, data *notifyModel.Notify, status int) error {
	data.Status = status

	filter := bson.D{{"id", data.Id}}
	update := bson.D{{"$set", bson.D{{"status", status}}}}
	err := biz.store.SolveNotify(ctx, filter, update)
	if err != nil {
		return err
	}
	if status == 0 {
		return nil
	}

	biz.pubsub.Publish(ctx, "AcpUser", pubsub.NewMessage(data))

	return nil
}
