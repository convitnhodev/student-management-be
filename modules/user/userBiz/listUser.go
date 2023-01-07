package userBiz

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/modules/user/userModel"

	"go.mongodb.org/mongo-driver/bson"
)

type ListUsersStore interface {
	ListUsersByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]userModel.User, error)
}

type listUsersBiz struct {
	store ListUsersStore
}

func NewListUsersBiz(store ListUsersStore) *listUsersBiz {
	return &listUsersBiz{store}
}

func (biz *listUsersBiz) ListUsers(ctx context.Context, filter interface{}, page *paging.Paging) ([]userModel.User, error) {
	// conditions := filter.(userModel.Filter)
	passingConditions := bson.D{{}}
	// if conditions.All == true {
	// 	passingConditions = bson.D{{}}
	// } else {
	// 	passingConditions = bson.D{{"acp", conditions.Acp}}
	// }

	data, err := biz.store.ListUsersByConditions(ctx, passingConditions, page)

	if err != nil {
		return nil, err
	}

	return data, nil
}
