package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationModel"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"
)

func ChangeAcpStudentAfterChangeNotify(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AcpStudent")

	store := studentStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			notify := msg.Data().(*notificationModel.Notification)
			var location string
			if notify.Location == "Class" {
				location = studentModel.Student_Class_Collection
			} else {
				location = studentModel.Student_Course_Collection
			}
			_ = store.UpdateStudent(ctx, bson.D{{"id", notify.Passive}}, bson.D{{"$set", bson.D{{"acp", true}}}}, location)
		}
	}()
}
