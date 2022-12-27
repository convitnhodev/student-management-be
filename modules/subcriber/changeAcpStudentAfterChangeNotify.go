package subcriber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notifyModel"
	"managerstudent/modules/student/studentStorage"
)

func ChangeAcpStudentAfterChangeNotify(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AcpStudent")

	store := studentStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			notify := msg.Data().(*notifyModel.Notify)
			var location string
			if notify.Location == "Class" {
				location = "student_class"
			} else {
				location = "student_course"
			}
			_ = store.UpdateStudent(ctx, bson.D{{"id", notify.Passive}}, bson.D{{"$set", bson.D{{"acp", true}}}}, location)
		}
	}()
}
