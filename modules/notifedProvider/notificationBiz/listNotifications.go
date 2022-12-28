package notificationBiz

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/modules/notifedProvider/notificationModel"
)

type ListNotificationsStore interface {
	ListNotificationsByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]notificationModel.Notification, error)
}

type getNotificationsBiz struct {
	store ListNotificationsStore
}

func NewListNotificationsBiz(store ListNotificationsStore) *getNotificationsBiz {
	return &getNotificationsBiz{store: store}
}

func (biz *getNotificationsBiz) ListNotifications(ctx context.Context, filter interface{}, page *paging.Paging) ([]notificationModel.Notification, error) {

	data, err := biz.store.ListNotificationsByConditions(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	return data, err
}
