package subjectBiz

import (
	"context"
	"managerstudent/modules/subject/subjectModel"
)

type ListSubjectStore interface {
	ListSubjects(ctx context.Context, filter interface{}) ([]subjectModel.Subject, error)
}



type listSubjectBiz struct {
	store ListSubjectStore
}


func NewListSubjectBiz(store ListSubjectStore) *listSubjectBiz {
	return &listSubjectBiz{store: store}
}


func (biz *listSubjectBiz) ListSubject(ctx context.Context, filter interface{}) ([]subjectModel.Subject, error) {
	result, err := biz.store.ListSubjects(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
