package notifyBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/paging"
	"managerstudent/modules/notifedProvider/notifyModel"
)

type ListNotificationsStore interface {
	ListNotificationsByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]notifyModel.Notify, error)
}

type getNotificationsBiz struct {
	store ListNotificationsStore
}

func NewListNotifcationsBiz(store ListNotificationsStore) *getNotificationsBiz {
	return &getNotificationsBiz{store: store}
}

func (biz *getNotificationsBiz) ListNotifications(ctx context.Context, filter interface{}, page *paging.Paging) ([]notifyModel.Notify, error) {

	data, err := biz.store.ListNotificationsByConditions(ctx, bson.M{"id": filter}, page)
	if err != nil {
		return nil, err
	}

	return data, err
}
