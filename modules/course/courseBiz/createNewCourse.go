package courseBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/course/courseModel"
)

type CreateCourseStore interface {
	CreateNewCourse(ctx context.Context, data *courseModel.Course) error
	FindCourse(ctx context.Context, conditions interface{}) (*courseModel.Course, error)
}

type createCourseBiz struct {
	store CreateCourseStore
}

func NewCreateCourseBiz(store CreateCourseStore) *createCourseBiz {
	return &createCourseBiz{store}
}

func (biz *createCourseBiz) CreateNewCourse(ctx context.Context, data *courseModel.Course) error {
	course, err := biz.store.FindCourse(ctx, bson.M{"course_id": data.Id})
	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if course != nil {
		managerLog.WarningLogger.Println("Course exist")
		return solveError.ErrEntityExisted("Course", nil)
	}

	if len(data.ListStudentId) < 1 {
		data.ListStudentId = make([]string, 0)
	}

	managerLog.InfoLogger.Println("Check student ok, can create currently class")
	if err := biz.store.CreateNewCourse(ctx, data); err != nil {
		managerLog.ErrorLogger.Println("Some thing error in storage class, may be from database")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Create course ok")
	return nil

}
