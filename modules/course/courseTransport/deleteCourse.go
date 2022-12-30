package courseTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/course/courseBiz"
	"managerstudent/modules/course/courseStorage"
)

func DeleteCourse(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string
		var ok bool

		id, ok = c.GetQuery("course_id")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}
		store := courseStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := courseBiz.NewDeleteCourseBiz(store, app.GetPubsub())
		err := biz.DeleteCourse(c.Request.Context(), &id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "success")
	}
}
