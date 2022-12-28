package studentTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentBiz"
	"managerstudent/modules/student/studentStorage"
)

func GetStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string
		var ok bool

		id, ok = c.GetQuery("id")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := studentBiz.NewGetStudent(store)
		data, err := biz.GetStudent(c.Request.Context(), &id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
