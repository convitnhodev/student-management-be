package classBiz

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/modules/class/classModel"
)

type ListClassesStore interface {
	ListClasses(ctx context.Context, conditions interface{}, page *paging.Paging) ([]classModel.Class, error)
}

type listClassesBiz struct {
	store ListClassesStore
}

func NewListClassesBiz(store ListClassesStore) *listClassesBiz {
	return &listClassesBiz{store}
}

func (biz *listClassesBiz) ListClasses(ctx context.Context, filter interface{}, page *paging.Paging) ([]classModel.Class, error) {
	data, err := biz.store.ListClasses(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	return data, nil
}
