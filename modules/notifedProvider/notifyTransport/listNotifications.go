package notifyTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notifyBiz"
	"managerstudent/modules/notifedProvider/notifyStorage"
)

func ListNotifications(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := notifyStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notifyBiz.NewListNotificationsBiz(store)
		data, err := biz.ListNotifications(c.Request.Context(), bson.D{}, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		page.Fullfill()
		c.JSON(200, customResponse.NewSuccessReponse(data, page, nil))
	}
}
