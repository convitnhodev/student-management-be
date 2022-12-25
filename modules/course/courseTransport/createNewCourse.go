package courseTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/course/courseBiz"
	"managerstudent/modules/course/courseModel"
	"managerstudent/modules/course/courseStorage"
)

func CreateNewCourse(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data courseModel.Course

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := courseStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := courseBiz.NewCreateCourseBiz(store)
		if err := biz.CreateNewCourse(c.Request.Context(), &data); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data.Id)
	}
}
