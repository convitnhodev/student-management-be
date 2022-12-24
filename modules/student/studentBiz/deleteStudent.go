package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type DeleteStudentStore interface {
	DeleteStudent(ctx context.Context, filter interface{}, location string) error
	FindStudent(ctx context.Context, conditions interface{}, location string) (*studentModel.Student, error)
}

type deleteStudentBiz struct {
	store DeleteStudentStore
}

func NewDeleteStudent(store DeleteStudentStore) *deleteStudentBiz {
	return &deleteStudentBiz{store: store}
}

func (biz *deleteStudentBiz) DeleteStudent(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteStudent(ctx, bson.M{"id": filter}, "student"); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create student ok")
	return nil
}

func (biz *deleteStudentBiz) DeleteStudentFromCourse(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteStudent(ctx, bson.M{"id": filter}, "student_course"); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create student ok")
	return nil
}

func (biz *deleteStudentBiz) DeleteStudentFromClass(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteStudent(ctx, bson.M{"id": filter}, "student_class"); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Create student ok")
	return nil
}
