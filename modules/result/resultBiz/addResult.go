package resultBiz

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

type AddResultStore interface {
	CreateListResult(ctx context.Context, data []resultModel.Result) error
}

type addResultBiz struct {
	store AddResultStore
}

func NewAddResultBiz(store AddResultStore) *addResultBiz {
	return &addResultBiz{store: store}
}

func (biz *addResultBiz) AddResult(ctx context.Context, data []resultModel.Result) error {
	err := biz.store.CreateListResult(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)

	}
	return nil
}
