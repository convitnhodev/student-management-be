package studentBiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/pubsub"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

type UserCreateStudentToCourseStore interface {
	CreateNewStudent(ctx context.Context, data interface{}, location string) error
	FindStudent(ctx context.Context, conditions interface{}, location string) (interface{}, error)
	UpdateStudent(ctx context.Context, conditions interface{}, value interface{}, location string) error
}

type userCreateStudentToCourseBiz struct {
	store  UserCreateStudentToCourseStore
	pubsub pubsub.Pubsub
}

func NewUserCreateStudentToCourseBiz(store UserCreateStudentToCourseStore, pubsub pubsub.Pubsub) *userCreateStudentToCourseBiz {
	return &userCreateStudentToCourseBiz{store: store, pubsub: pubsub}
}

func (biz *userCreateStudentToCourseBiz) UserCreateNewStudentToCourse(ctx context.Context, data *studentModel.StudentAndCourse) error {

	student, err := biz.store.FindStudent(ctx, bson.M{"id": data.StudentId}, studentModel.StudentCollectionFullInfo)

	if err != nil {
		if err.Error() != solveError.RecordNotFound {
			managerLog.ErrorLogger.Println("Some thing error in storage student, may be from database")
			return solveError.ErrDB(err)
		}
	}

	if student == nil {
		managerLog.WarningLogger.Println("Student is not exist")
		return solveError.ErrEntityNotExisted("Student", nil)
	}

	list_courses := make([]string, 0)
	for _, value := range data.Courses {
		list_courses = append(list_courses, value)
	}

	filter := bson.D{{"id", data.StudentId}}
	update := bson.M{"$push": bson.M{"list_course_id": bson.M{"$each": list_courses}}}
	if err := biz.store.UpdateStudent(ctx, filter, update, studentModel.StudentCollection); err != nil {
		return err

	}

	_ = biz.pubsub.Publish(ctx, "AddStudentToCourse", pubsub.NewMessage(data))

	return nil
}
