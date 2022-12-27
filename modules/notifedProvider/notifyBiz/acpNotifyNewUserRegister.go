package notifyBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notifyModel"
)

type AcpNotifyUserRegisterStore interface {
	SolveNotify(ctx context.Context, conditions interface{}, value interface{}) error
}

type acpNotifyUserRegisterBiz struct {
	store  AcpNotifyUserRegisterStore
	pubsub pubsub.Pubsub
}

func NewAcpNotifyUserRegisterBiz(store AcpNotifyUserRegisterStore, pubsub pubsub.Pubsub) *acpNotifyUserRegisterBiz {
	return &acpNotifyUserRegisterBiz{store: store, pubsub: pubsub}
}

func (biz *acpNotifyUserRegisterBiz) AcpNotifyUserRegister(ctx context.Context, data *notifyModel.Notify, status int) error {
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
