package studentBiz

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"

	"go.mongodb.org/mongo-driver/bson"
)

type GetStudentStore interface {
	FindStudent(ctx context.Context, conditions interface{}) (*studentModel.Student, error)
}

type getStudentBiz struct {
	store GetStudentStore
}

func NewGetStudent(store GetStudentStore) *getStudentBiz {
	return &getStudentBiz{store: store}
}

func (biz *getStudentBiz) GetStudent(ctx context.Context, filter interface{}) (*studentModel.Student, error) {
	data, err := biz.store.FindStudent(ctx, bson.M{"id": filter})
	if err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return nil, solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Get student ok")
	return data, nil
}
