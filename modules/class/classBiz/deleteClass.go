package classBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

type DeleteClassStore interface {
	DeleteClass(ctx context.Context, conditions interface{}) error
}

type deleteClassBiz struct {
	store  DeleteClassStore
	pubsub pubsub.Pubsub
}

func NewDeleteClassBiz(store DeleteClassStore, pubsub pubsub.Pubsub) *deleteClassBiz {
	return &deleteClassBiz{store, pubsub}
}

func (biz *deleteClassBiz) DeleteClass(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteClass(ctx, bson.M{"class_id": filter}); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage class, may be from database")
		return solveError.ErrDB(err)
	}

	biz.pubsub.Publish(ctx, "DeleteClass", pubsub.NewMessage(filter))
	return nil
}
