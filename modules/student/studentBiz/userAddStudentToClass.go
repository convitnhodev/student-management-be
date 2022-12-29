package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type UserCreateStudentToClassStore interface {
	CreateNewStudent(ctx context.Context, data interface{}, location string) error
	FindStudent(ctx context.Context, conditions interface{}, location string) (interface{}, error)
	UpdateStudent(ctx context.Context, conditions interface{}, value interface{}, location string) error
}

type userCreateStudentToClassBiz struct {
	store  UserCreateStudentToClassStore
	pubsub pubsub.Pubsub
}

func NewUserCreateStudentToClassBiz(store UserCreateStudentToClassStore, pubsub pubsub.Pubsub) *userCreateStudentToClassBiz {
	return &userCreateStudentToClassBiz{store: store, pubsub: pubsub}
}

func (biz *userCreateStudentToClassBiz) UserCreateNewStudentToClass(ctx context.Context, data *studentModel.StudentAndClass) error {

	student, err := biz.store.FindStudent(ctx, bson.M{"id": data.StudentId}, studentModel.StudentCollectionFullInfo)

	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if student == nil {
		managerLog.WarningLogger.Println("Student is not exist")
		return solveError.ErrEntityNotExisted("Student", nil)
	}

	filter := bson.D{{"id", data.StudentId}}
	update := bson.D{{"$set", bson.D{{"class_id", data.ClassId}}}}
	if err := biz.store.UpdateStudent(ctx, filter, update, studentModel.StudentCollection); err != nil {
		return err
	}

	_ = biz.pubsub.Publish(ctx, "AddStudentToClass", pubsub.NewMessage(data))


	return nil
}
