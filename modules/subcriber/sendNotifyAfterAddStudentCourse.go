package subcriber

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationModel"
	"managerstudent/modules/notifedProvider/notificationStorage"
)

func SendNotifyAfterAddStudentToCourse(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AddStudentToCourseNotify")

	store := notificationStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {
			msg := <-c
			notify := msg.Data().(notificationModel.Notification)
			_ = store.CreateNewNotification(ctx, &notify)
		}
	}()
}
