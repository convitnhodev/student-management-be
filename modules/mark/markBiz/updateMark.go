package markBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/mark/markModel"
)

type UpdateResultStore interface {
	UpdateResult(ctx context.Context, conditions interface{}, data markModel.Result) error
}

type updateResultBiz struct {
	store UpdateResultStore
}

func NewUpdateResultBiz(store UpdateResultStore) *updateResultBiz {
	return &updateResultBiz{store: store}
}

func (biz *updateResultBiz) UpdateResult(ctx context.Context, data markModel.Result) error {
	err := biz.store.UpdateResult(ctx, bson.M{"id_student": data.IdStudent}, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)

	}
	return nil
}
