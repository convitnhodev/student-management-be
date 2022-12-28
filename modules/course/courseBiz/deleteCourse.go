package courseBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

type DeleteCourseStore interface {
	DeleteCourse(ctx context.Context, conditions interface{}) error
}

type deleteCourseBiz struct {
	store  DeleteCourseStore
	pubsub pubsub.Pubsub
}

func NewDeleteCourseBiz(store DeleteCourseStore, pubsub pubsub.Pubsub) *deleteCourseBiz {
	return &deleteCourseBiz{store, pubsub}
}

func (biz *deleteCourseBiz) DeleteCourse(ctx context.Context, filter interface{}) error {
	if err := biz.store.DeleteCourse(ctx, bson.M{"course_id": filter}); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage course, may be from database")
		return solveError.ErrDB(err)
	}
	biz.pubsub.Publish(ctx, "DeleteCourse", pubsub.NewMessage(filter))
	return nil
}
