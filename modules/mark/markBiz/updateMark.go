package markBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type UpdateResultStore interface {
	UpdateResult(ctx context.Context, data []studentModel.Result, conditions interface{}) error
}

type updateResultBiz struct {
	store UpdateResultStore
}

func NewUpdateResultBiz(store UpdateResultStore) *updateResultBiz {
	return &updateResultBiz{store}
}

func (biz *updateResultBiz) UpdateNewResult(ctx context.Context, data []studentModel.Result) error {

	conditions := bson.D{{""}}
	err := biz.store.UpdateResult(ctx, data, filters)

	if err != nil {
		managerLog.ErrorLogger.Println("Can not update result")
		return err
	}

	managerLog.InfoLogger.Println("Update ok")
	return nil

}
