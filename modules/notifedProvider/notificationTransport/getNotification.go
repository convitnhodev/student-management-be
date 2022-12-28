package notificationTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationBiz"
	"managerstudent/modules/notifedProvider/notificationStorage"
)

func GetNotification(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ok bool

		tmp, ok := c.GetQuery("id")
		//id, _ = strconv.Atoi(tmp)
		id, err := primitive.ObjectIDFromHex(tmp)
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		store := notificationStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notificationBiz.NewGetNotifyBiz(store, app.GetPubsub())
		data, err := biz.GetNotification(c, id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data)
	}
}
