package notifyTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/notifedProvider/notifyBiz"
	"managerstudent/modules/notifedProvider/notifyStorage"
	"strconv"
)

func GetNotify(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id int
		var ok bool

		tmp, ok := c.GetQuery("id")
		id, _ = strconv.Atoi(tmp)
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		//
		store := notifyStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := notifyBiz.NewGetNotifyBiz(store, app.GetPubsub())
		data, err := biz.GetNotify(c, id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data)
	}
}
