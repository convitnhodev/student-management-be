package courseBiz

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/modules/course/courseModel"
)

type ListCoursesStore interface {
	ListCourses(ctx context.Context, conditions interface{}, page *paging.Paging) ([]courseModel.Course, error)
}

type listCoursesBiz struct {
	store ListCoursesStore
}

func NewListCoursesBiz(store ListCoursesStore) *listCoursesBiz {
	return &listCoursesBiz{store}
}

func (biz *listCoursesBiz) ListCourses(ctx context.Context, filter interface{}, page *paging.Paging) ([]courseModel.Course, error) {
	data, err := biz.store.ListCourses(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	return data, nil
}
