package classTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/class/classStorage"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListStudents(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		classId := c.Request.URL.Query().Get("class_id")

		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		data, err := store.FindClass(c.Request.Context(), bson.M{"class_id": classId})
		if err != nil {
			c.JSON(400, err)
			return
		}

		var students []studentModel.Student
		studentStore := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		for _, v := range data.ListStudentId {
			student, _ := studentStore.FindStudent(c.Request.Context(), bson.M{"id": v})
			students = append(students, *student)
		}

		c.JSON(200, customResponse.SimpleSuccessReponse(students))
	}
}
