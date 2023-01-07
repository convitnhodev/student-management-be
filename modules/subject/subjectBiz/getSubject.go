package subjectBiz

import (
	"context"
	"managerstudent/modules/subject/subjectModel"
)

type GetSubjectStore interface {
	GetSubject(ctx context.Context, filter interface{}) (*subjectModel.Subject, error)
}

type getSubjectBiz struct {
	store GetSubjectStore
}

func NewGetSubjectBiz(store GetSubjectStore) *getSubjectBiz {
	return &getSubjectBiz{store: store}
}

func (biz *getSubjectBiz) GetSubject(ctx context.Context, filter interface{}) (*subjectModel.Subject, error) {
	data, err := biz.store.GetSubject(ctx, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}
