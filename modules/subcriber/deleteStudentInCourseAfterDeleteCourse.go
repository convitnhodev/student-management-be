package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentStorage"
)

func DeleteStudentInCourseAfterDeleteCourse(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "deleteCourse")

	store := studentStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			CourseId := msg.Data().(*string)
			_ = store.DeleteManyStudent(ctx, bson.M{"course_id": CourseId}, "student_course")
		}
	}()
}
