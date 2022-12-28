package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"
)

func DeleteStudentInCourseAfterDeleteCourse(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "DeleteCourse")

	store := studentStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			CourseId := msg.Data().(*string)
			_ = store.DeleteManyStudent(ctx, bson.M{"course_id": CourseId}, studentModel.Student_Course_Collection)
		}
	}()
}
