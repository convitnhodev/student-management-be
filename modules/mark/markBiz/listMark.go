package markBiz

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/mark/markModel"
	"managerstudent/modules/student/studentModel"
)

type ListResultStore interface {
	ListResultByConditions(ctx context.Context, conditions interface{}) (*studentModel.Result, error)
}

type listResultBiz struct {
	store ListResultStore
}

func NewListMarkBiz(store ListResultStore) *listResultBiz {
	return &listResultBiz{store: store}
}

func (biz *listResultBiz) ListResult(ctx context.Context, filters *markModel.Filter) (*studentModel.Result, error) {
	result, err := biz.store.ListResultByConditions(ctx, filters)
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage mark, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}

	return result, nil
}
