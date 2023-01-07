package classTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classBiz"
	"managerstudent/modules/class/classStorage"

	"github.com/gin-gonic/gin"
)

func GetClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var id string
		var ok bool

		id, ok = c.GetQuery("class_id")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := classBiz.NewGetClassesBiz(store)
		data, err := biz.GetClass(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
