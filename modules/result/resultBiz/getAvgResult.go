package resultBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type CountResultStore interface {
	ListResultByConditions(ctx context.Context, conditions interface{}) ([]studentModel.Result, error)
}

type countResultBiz struct {
	store CountResultStore
}

func NewCountAvgMarkBiz(store CountResultStore) *countResultBiz {
	return &countResultBiz{store: store}
}

func (biz *countResultBiz) CountResult(ctx context.Context, idStudent interface{}) (*float64, error) {
	result, err := biz.store.ListResultByConditions(ctx, bson.M{"id_student": idStudent})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage mark, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}
	var sum float64
	for _, value := range result {
		sum += value.Average
	}
	average := sum / float64(len(result))

	return &average, nil
}
