package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classStorage"
	"managerstudent/modules/student/studentModel"
)

func AddStudentIdAfterAddStudentToClass(appCtx component.AppContext, ctx context.Context)  {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AddStudentToClass")

	store := classStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {
			msg := <-c
			data := msg.Data().(*studentModel.StudentAndClass)
			filter := bson.D{{"class_id", data.ClassId}}
			update := bson.M{"$push": bson.M{"list_student_id":  data.StudentId}, "$inc": bson.M{"total": 1}}
			_ = store.UpdateClass(ctx, filter, update)
		}
	}()
}


