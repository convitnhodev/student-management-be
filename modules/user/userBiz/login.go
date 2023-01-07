package userBiz

import (
	"context"
	"errors"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/component/hasher"
	"managerstudent/component/tokenProvider"
	"managerstudent/modules/user/userModel"

	"go.mongodb.org/mongo-driver/bson"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenProvider.Provider
	hasher        hasher.HasherInfo
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenProvider.Provider, hasher hasher.HasherInfo, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *userModel.UserLogin) (*tokenProvider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, bson.M{"username": data.UserName})
	if err != nil {
		return nil, solveError.ErrEntityNotExisted("User", nil)
	}

	if user == nil {
		return nil, solveError.ErrEntityNotExisted("User", nil)
	}

	if !user.Acp {
		return nil, solveError.ErrEntityNotExisted("User", errors.New("User is not acp"))
	}

	if user.Password != biz.hasher.HashMd5(user.Salt+data.Password+user.Salt) {
		return nil, solveError.ErrInvalidLogin(errors.New("info is invalid"))
	}

	payload := tokenProvider.TokenPayload{
		UserName: user.UserName,
		Role:     user.Role,
	}

	token, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}

	return token, nil
}
