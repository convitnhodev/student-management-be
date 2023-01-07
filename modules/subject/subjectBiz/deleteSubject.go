package subjectBiz

import "context"

type DeleteSubjectStore interface {
	DeleteSubject(ctx context.Context, conditions interface{}) error
}



type deleteSubjectBiz struct {
	store DeleteSubjectStore
}


func NewDeleteSubjectBiz (store DeleteSubjectStore) *deleteSubjectBiz {
	return &deleteSubjectBiz{store: store}
}


func (biz *deleteSubjectBiz) DeleteSubject(ctx context.Context, conditions interface{}) error {
	if err := biz.store.DeleteSubject(ctx, conditions); err != nil {
		return err
	}
	return nil
}



