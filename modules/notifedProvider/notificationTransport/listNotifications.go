package notificationTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationBiz"
	"managerstudent/modules/notifedProvider/notificationStorage"
)

func ListNotifications(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := notificationStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notificationBiz.NewListNotificationsBiz(store)
		data, err := biz.ListNotifications(c.Request.Context(), bson.D{}, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		page.Fullfill()
		c.JSON(200, customResponse.NewSuccessReponse(data, page, nil))
	}
}
