package classBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/class/classModel"
)

type CreateClassStore interface {
	CreateNewClass(ctx context.Context, data *classModel.Class) error
	FindClass(ctx context.Context, conditions interface{}) (*classModel.Class, error)
}

type createClassBiz struct {
	store CreateClassStore
}

func NewCreateClassBiz(store CreateClassStore) *createClassBiz {
	return &createClassBiz{store}
}

func (biz *createClassBiz) CreateNewClass(ctx context.Context, data *classModel.Class) error {
	class, err := biz.store.FindClass(ctx, bson.M{"class_id": data.Id})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if class != nil {
		managerLog.WarningLogger.Println("Class exist")
		return solveError.ErrEntityExisted("Class", nil)
	}

	managerLog.InfoLogger.Println("Check student ok, can create currently class")
	if err := biz.store.CreateNewClass(ctx, data); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage class, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create class ok")
	return nil

}




