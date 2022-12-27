package notifyTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notifyBiz"
	"managerstudent/modules/notifedProvider/notifyModel"
	"managerstudent/modules/notifedProvider/notifyStorage"
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

		var data notifyModel.Notify
		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		//
		store := notifyStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notifyBiz.NewAcpNotifyRequestAddStudentBiz(store, app.GetPubsub())
		if err := biz.AcpNotifyRequestAddStudent(c.Request.Context(), &data, status); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data)
	}
}
