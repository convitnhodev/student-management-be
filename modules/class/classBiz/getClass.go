package classBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/modules/class/classModel"
)

type GetClassesStore interface {
	FindClass(ctx context.Context, conditions interface{}) (*classModel.Class, error)
}

type getClassesBiz struct {
	store GetClassesStore
}

func NewGetClassesBiz(store GetClassesStore) *getClassesBiz {
	return &getClassesBiz{store}
}

func (biz *getClassesBiz) GetClass(ctx context.Context, filter interface{}) (*classModel.Class, error) {
	data, err := biz.store.FindClass(ctx, bson.D{{"class_id", filter}})
	if err != nil {
		return nil, err
	}
	return data, nil
}
