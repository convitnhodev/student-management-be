package resultBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/student/studentModel"
)

type AddResultStore interface {
	CreateListResult(ctx context.Context, data []resultModel.Result) error
}

type CheckCourseStore interface {
	FindStudent(ctx context.Context, conditions interface{}, location string) (interface{}, error)
}

type addResultBiz struct {
	store       AddResultStore
	checkCourse CheckCourseStore
}

func NewAddResultBiz(store AddResultStore, courseStore CheckCourseStore) *addResultBiz {
	return &addResultBiz{store: store, checkCourse: courseStore}
}

func (biz *addResultBiz) AddResult(ctx context.Context, data resultModel.Result) error {
	student, err := biz.checkCourse.FindStudent(ctx, bson.M{"id": data.StudentId}, studentModel.StudentCollectionFullInfo)

	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if student != nil {
		managerLog.WarningLogger.Println("Student exist")
		return solveError.ErrEntityExisted("Student", nil)
	}

	if err := biz.store.CreateNewStudent(ctx, data, studentModel.StudentCollectionFullInfo); err != nil {
		return nil, err
	}

	err := biz.store.CreateListResult(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)

	}
	return nil
}
