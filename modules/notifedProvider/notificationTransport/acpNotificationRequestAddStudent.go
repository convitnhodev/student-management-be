package notificationTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notificationBiz"
	"managerstudent/modules/notifedProvider/notificationModel"
	"managerstudent/modules/notifedProvider/notificationStorage"
	"strconv"
)

func AdminAcpNotifyRequestAddStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var status int
		var ok bool

		tmp, ok := c.GetQuery("status")
		status, _ = strconv.Atoi(tmp)
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		var data notificationModel.Notification
		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		//
		store := notificationStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notificationBiz.NewAcpNotificationRequestAddStudentBiz(store, app.GetPubsub())
		if err := biz.AcpNotificationRequestAddStudent(c.Request.Context(), &data, status); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data)
	}
}
