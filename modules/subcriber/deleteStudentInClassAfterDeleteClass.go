package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"
)

func DeleteStudentInClassAfterDeleteClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "DeleteClass")

	store := studentStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			ClassId := msg.Data().(*string)
			_ = store.DeleteManyStudent(ctx, bson.M{"class_id": ClassId}, studentModel.Student_Class_Collection)
		}
	}()
}
