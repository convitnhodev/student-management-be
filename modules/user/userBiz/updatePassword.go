package userBiz

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/hasher"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"
)

type UpdatePasswordStore interface {
	UpdateUser(ctx context.Context, conditions interface{}, data interface{}) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type updatePasswordBusiness struct {
	store  UpdateUserStore
	hasher hasher.HasherInfo
}

func NewUpdatepasswordBusiness(store UpdateUserStore, hasher hasher.HasherInfo) *updatePasswordBusiness {
	return &updatePasswordBusiness{store: store, hasher: hasher}
}

func (biz *updatePasswordBusiness) UpdatePasswordOfUser(ctx context.Context, conditions interface{}, data *userModel.UpdatePassWord) error {

	user, err := biz.store.FindUser(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
			return solveError.ErrDB(err)
		}
	}
	if user == nil {
		managerLog.WarningLogger.Println("User's not new")
		return solveError.ErrEntityExisted("User is not exist", nil)
	}
	fmt.Print(biz.hasher.HashMd5(user.Salt + data.Password + user.Salt))

	managerLog.InfoLogger.Println("Check user ok, can create currently user")

	if biz.hasher.HashMd5(user.Salt+data.Password+user.Salt) != user.Password {
		return solveError.ErrInvalidCurrentPassword()
	}

	pass := bson.M{"pass_word": biz.hasher.HashMd5(user.Salt + data.NewPassword + user.Salt)}
	if err := biz.store.UpdateUser(ctx, conditions, pass); err != nil {
		return err
	}
	return nil
}
