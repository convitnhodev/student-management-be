package subcriber

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/course/courseStorage"
	"managerstudent/modules/student/studentModel"
)

func AddStudentIdAfterAddStudentToCourse(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AddStudentToCourse")

	store := courseStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {
			msg := <-c
			data := msg.Data().(*studentModel.StudentAndCourse)
			coursesid := data.Courses

			for _, value := range coursesid {
				fmt.Print(value)
				fmt.Println("hello")
				filter := bson.D{{"course_id", value}}
				update := bson.M{"$push": bson.M{"list_student_id": data.StudentId}, "$inc": bson.M{"total": 1}}
				_ = store.UpdateCourse(ctx, filter, update)
			}
		}
	}()
}
