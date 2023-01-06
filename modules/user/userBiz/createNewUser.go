package userBiz

import (
	"context"
	"fmt"
	"managerstudent/common/pubsub"
	generatesalt "managerstudent/common/salt"
	"managerstudent/common/solveError"
	"managerstudent/component/hasher"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notificationModel"
	"managerstudent/modules/user/userModel"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *userModel.User) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type createUserBiz struct {
	store  CreateUserStore
	hasher hasher.HasherInfo
	pubsub pubsub.Pubsub
}

func NewCreateUserBiz(store CreateUserStore, hasher hasher.HasherInfo, pubsub pubsub.Pubsub) *createUserBiz {
	return &createUserBiz{store, hasher, pubsub}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *userModel.User) error {
	user, err := biz.store.FindUser(ctx, bson.M{"username": data.UserName})
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
	salt := generatesalt.GenSalt(50)
	data.Salt = salt
	data.Acp = false
	data.Password = biz.hasher.HashMd5(salt + data.Password + salt)
	// data.Role = _const.Role(0)
	fmt.Println(data)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage user, may be from database")
		return solveError.ErrDB(err)
	}

	notify := notificationModel.Notification{
		Content: fmt.Sprint(data.UserName, " yeu cau dang ki tai khoan"),
		Agent:   data.UserName,
		Seen:    false,
		Time:    time.Now(),
	}
	biz.pubsub.Publish(ctx, "RegisterNotify", pubsub.NewMessage(notify))

	managerLog.InfoLogger.Println("Create user ok")
	return nil
}
