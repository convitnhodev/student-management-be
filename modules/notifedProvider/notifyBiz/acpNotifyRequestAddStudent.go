package notifyBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/modules/notifedProvider/notifyModel"
)

type AcpNotifyRequestAddStudentStore interface {
	SolveNotify(ctx context.Context, conditions interface{}, value interface{}) error
}

type acpNotifyRequestAddStudentBiz struct {
	store  AcpNotifyRequestAddStudentStore
	pubsub pubsub.Pubsub
}

func NewAcpNotifyRequestAddStudentBiz(store AcpNotifyRequestAddStudentStore, pubsub pubsub.Pubsub) *acpNotifyRequestAddStudentBiz {
	return &acpNotifyRequestAddStudentBiz{store: store, pubsub: pubsub}
}

func (biz *acpNotifyRequestAddStudentBiz) AcpNotifyRequestAddStudent(ctx context.Context, data *notifyModel.Notify, status int) error {
	data.Status = status

	filter := bson.D{{"id", data.Id}}
	update := bson.D{{"$set", bson.D{{"status", status}, {"seen", true}}}}
	err := biz.store.SolveNotify(ctx, filter, update)
	if err != nil {
		return err
	}
	if status == 0 {
		return nil
	}

	biz.pubsub.Publish(ctx, "AcpStudent", pubsub.NewMessage(data))

	return nil
}
