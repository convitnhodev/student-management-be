package studentTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentBiz"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"
)

func AddStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student studentModel.Student
		if err := c.ShouldBind(&student); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := studentBiz.NewAddStudentBiz(store)
		if err := biz.AddStudentTo(c.Request.Context(), &student); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "ok")
	}
}

func AddStudentToClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student studentModel.Student
		if err := c.ShouldBind(&student); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := studentBiz.NewAddStudentBiz(store)
		if err := biz.AddStudentToClass(c.Request.Context(), &student); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "ok")
	}
}

func AddStudentToCourse(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student studentModel.Student
		if err := c.ShouldBind(&student); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := studentBiz.NewAddStudentBiz(store)
		if err := biz.AddStudentToCourse(c.Request.Context(), &student); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "ok")
	}
}
