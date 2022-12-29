package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type AdminCreateStudentStore interface {
	CreateNewStudent(ctx context.Context, data interface{}, location string) error
	FindStudent(ctx context.Context, conditions interface{}, location string) (interface{}, error)
}

type adminCreateStudentBiz struct {
	store  AdminCreateStudentStore
	pubsub pubsub.Pubsub
}

func NewAdminCreateStudentBiz(store AdminCreateStudentStore, pubsub pubsub.Pubsub) *adminCreateStudentBiz {
	return &adminCreateStudentBiz{store: store, pubsub: pubsub}
}

func (biz *adminCreateStudentBiz) CreateNewStudent(ctx context.Context, data *studentModel.FullInfoStudent) (*studentModel.FullInfoStudent, error) {

	student, err := biz.store.FindStudent(ctx, bson.M{"id": data.Id}, studentModel.StudentCollectionFullInfo)

	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return nil, solveError.ErrDB(err)
		}
	}

	if student != nil {
		managerLog.WarningLogger.Println("Student exist")
		return nil, solveError.ErrEntityExisted("Student", nil)
	}

	if err := biz.store.CreateNewStudent(ctx, data, studentModel.StudentCollectionFullInfo); err != nil {
		return nil, err
	}

	addingData := studentModel.Student{
		Id:      data.Id,
		Results: make([]studentModel.Result, 0),
	}
	if err := biz.store.CreateNewStudent(ctx, addingData, studentModel.StudentCollection); err != nil {
		return nil, err
	}

	return data, nil
}
