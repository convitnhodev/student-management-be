package resultBiz

import (
	"context"
	"managerstudent/modules/result/resultModel"
)

type AddResultStore interface {
	CreateResult(ctx context.Context, data resultModel.Result) error
}

type addResultBiz struct {
	store AddResultStore
}

func NewAddResultBiz(store AddResultStore) *addResultBiz {
	return &addResultBiz{store: store}
}

func (biz *addResultBiz) CreateOrUpdateResult(ctx context.Context, data resultModel.Result) error {
	data.CalculateAverage()
	if err := biz.store.CreateResult(ctx, data); err != nil {
		return err
	}
	return nil
}
