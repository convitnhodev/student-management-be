package courseTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classModel"
	"managerstudent/modules/course/courseBiz"
	"managerstudent/modules/course/courseStorage"
)

func ListCourses(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		var filter classModel.Filter
		if err := c.BindJSON(&filter); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := courseStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := courseBiz.NewListCoursesBiz(store)
		data, err := biz.ListCourses(c.Request.Context(), filter, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data)
	}
}
