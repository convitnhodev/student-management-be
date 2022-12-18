package userBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/hasher"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *userModel.User) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type createUserBiz struct {
	store  CreateUserStore
	hasher hasher.HasherInfo
}

func NewCreateUserBiz(store CreateUserStore, hasher hasher.HasherInfo) *createUserBiz {
	return &createUserBiz{store, hasher}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *userModel.User) error {
	user, err := biz.store.FindUser(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
			return solveError.ErrDB(err)
		}
	}
	if user != nil {
		managerLog.WarningLogger.Println("User's not new")
		return solveError.ErrEntityExisted("User", nil)
	}

	managerLog.InfoLogger.Println("Check user ok, can create currently user")

	data.Password = biz.hasher.HashMd5(data.Password)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create user ok")
	return nil
}
