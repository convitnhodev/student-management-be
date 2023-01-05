package subcriber

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationModel"
	"managerstudent/modules/user/userStorage"

	"go.mongodb.org/mongo-driver/bson"
)

func ChangeAcpUserAfterChangeNotify(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, "AcpUser")

	store := userStorage.NewMongoStore(appCtx.GetNewDataMongoDB())

	go func() {
		defer solveError.AppRecover()
		for {

			msg := <-c
			notify := msg.Data().(*notificationModel.Notification)
			_ = store.UpdateUser(ctx, bson.D{{"username", notify.Agent}}, bson.D{{"$set", bson.D{{"acp", true}}}})
		}
	}()
}
