package markTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/mark/markBiz"
	"managerstudent/modules/mark/markModel"
	"managerstudent/modules/mark/markStorage"
)

func ListResultByIdStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter markModel.Filter
		var ok bool

		filter.IdStudent, ok = c.GetQuery("id_student")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		store := markStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := markBiz.NewListMarkBiz(store)
		data, err := biz.ListResultByIdStudent(c.Request.Context(), &filter.IdStudent)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, data)
	}
}

func ListResultByIdClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter markModel.Filter
		var ok bool

		filter.IdClass, ok = c.GetQuery("id_class")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		store := markStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := markBiz.NewListMarkBiz(store)
		data, err := biz.ListResultByIdClass(c.Request.Context(), &filter.IdClass)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, data)
	}
}

func ListResultByIdCourse(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter markModel.Filter
		var ok bool

		filter.IdCourse, ok = c.GetQuery("id_course")
		if !ok {
			panic(solveError.ErrInvalidRequest(nil))
		}

		store := markStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := markBiz.NewListMarkBiz(store)
		data, err := biz.ListResultByIdCourse(c.Request.Context(), &filter.IdCourse)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, data)
	}
}
