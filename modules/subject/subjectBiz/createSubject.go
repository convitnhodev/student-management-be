package subjectBiz

import (
	"context"
	"managerstudent/modules/subject/subjectModel"
)

type CreateSubjectStore interface {
	CreateNewSubject(ctx context.Context, data *subjectModel.Subject) error
}


type createSubjectBiz struct {
	store CreateSubjectStore
}

func NewCreateSubjectBiz(store CreateSubjectStore) *createSubjectBiz {
	return &createSubjectBiz{store: store}
}


func (biz *createSubjectBiz) CreateSubject(ctx context.Context, data *subjectModel.Subject) error {
	if err := biz.store.CreateNewSubject(ctx, data); err != nil {
		return err
	}
	return nil
}





