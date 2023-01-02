package resultBiz

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

type DeleteResultStore interface {
	DeleteResult(ctx context.Context, filter interface{}) error
}

type deleteResultBiz struct {
	store DeleteResultStore
}

func NewDeleteResultBiz(store DeleteResultStore) *deleteResultBiz {
	return &deleteResultBiz{store: store}
}

func (biz *deleteResultBiz) DeleteResult(ctx context.Context, filter interface{}) error {
	err := biz.store.DeleteResult(ctx, filter)
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)

	}
	return nil
}
