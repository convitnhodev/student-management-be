package resultBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type ListResultStore interface {
	ListResultByConditions(ctx context.Context, conditions interface{}) ([]studentModel.Result, error)
}

type listResultBiz struct {
	store ListResultStore
}

func NewListMarkBiz(store ListResultStore) *listResultBiz {
	return &listResultBiz{store: store}
}

func (biz *listResultBiz) ListResultByIdStudent(ctx context.Context, conditions interface{}) ([]studentModel.Result, error) {
	result, err := biz.store.ListResultByConditions(ctx, bson.M{"id_student": conditions})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage mark, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}

	return result, nil
}

func (biz *listResultBiz) ListResultByIdClass(ctx context.Context, conditions interface{}) ([]studentModel.Result, error) {
	result, err := biz.store.ListResultByConditions(ctx, bson.M{"id_class": conditions})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage mark, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}

	return result, nil
}

func (biz *listResultBiz) ListResultByIdCourse(ctx context.Context, conditions interface{}) ([]studentModel.Result, error) {
	result, err := biz.store.ListResultByConditions(ctx, bson.M{"id_course": conditions})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage mark, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}

	return result, nil
}
