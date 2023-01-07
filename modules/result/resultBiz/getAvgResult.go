package resultBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/paging"
	"managerstudent/modules/result/resultModel"
)

type CountResultStore interface {
	ListResultByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]resultModel.Result, error)
}

type countResultBiz struct {
	store CountResultStore
}

func NewCountAvgMarkBiz(store CountResultStore) *countResultBiz {
	return &countResultBiz{store: store}
}

func (biz *countResultBiz) CountResult(ctx context.Context, idStudent interface{}) (*float64, error) {
	result, err := biz.store.ListResultByConditions(ctx, bson.M{"id_student": idStudent}, nil)
	if err != nil {
		return nil, err
	}
	var sum float64
	for _, value := range result {
		sum += value.Average
	}
	average := sum / float64(len(result))

	return &average, nil
}
