package resultBiz

import (
	"context"
	"managerstudent/modules/result/resultModel"
)

type UpdateResultStore interface {
	UpdateResult(ctx context.Context, filter interface{}, data resultModel.Result) error
}

type updateResultBiz struct {
	store UpdateResultStore
}

func NewUpdateResultBiz(store UpdateResultStore) *updateResultBiz {
	return &updateResultBiz{store: store}
}

func (biz *updateResultBiz) UpdateResult(ctx context.Context, filter interface{}, data resultModel.Result) error {
	data.CalculateAverage()
	if err := biz.store.UpdateResult(ctx, filter, data); err != nil {
		return err
	}
	return nil
}
