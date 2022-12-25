package courseBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

type DeleteCourseStore interface {
	DeleteCourse(ctx context.Context, conditions interface{}) error
}

type deleteCourseBiz struct {
	store DeleteCourseStore
}

func NewDeleteCourseBiz(store DeleteCourseStore) *deleteCourseBiz {
	return &deleteCourseBiz{store}
}

func (biz *deleteCourseBiz) DeleteCourse(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteCourse(ctx, bson.M{"course_id": filter}); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage course, may be from database")
		return solveError.ErrDB(err)
	}
	return nil
}
