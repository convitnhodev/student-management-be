package resultBiz

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/modules/result/resultModel"
)

type ListResultStore interface {
	ListResultByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]resultModel.Result, error)
}

type listResultBiz struct {
	store ListResultStore
}

func NewListMarkBiz(store ListResultStore) *listResultBiz {
	return &listResultBiz{store: store}
}

func (biz *listResultBiz) ListResult(ctx context.Context, filter interface{}, page *paging.Paging) ([]resultModel.Result, error) {
	result, err := biz.store.ListResultByConditions(ctx, filter, page)
	if err != nil {
		return nil, err

	}
	return result, nil
}
