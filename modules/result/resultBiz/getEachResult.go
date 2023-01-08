package resultBiz

import (
	"context"
	"managerstudent/modules/result/resultModel"
)

type GetResultStore interface {
	GetResult(ctx context.Context, conditions interface{}) (*resultModel.Result, error)
}

type getResultBiz struct {
	store GetResultStore
}

func NewGetMarkBiz(store GetResultStore) *getResultBiz {
	return &getResultBiz{store: store}
}

func (biz *getResultBiz) GetEachResult(ctx context.Context, filter interface{}) (*resultModel.Result, error) {
	result, err := biz.store.GetResult(ctx, filter)
	if err != nil {
		return nil, err

	}
	return result, nil
}
