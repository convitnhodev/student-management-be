package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"managerstudent/common/setupDatabase"
	"managerstudent/component"
	"managerstudent/component/managerLog"
	"managerstudent/modules/mark/markTransport"
	"managerstudent/modules/student/studentTransport"
	"managerstudent/modules/user/userTransport"
)

func main() {
	managerLog.InitLogs()
	db := setupDatabase.InitMongoDB()
	fmt.Println(db)
	if err := runService(db, nil); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *mongo.Client, redis *redis.Client) error {
	r := gin.Default()
	time := component.TimeJWT{60 * 60 * 24 * 2, 60 * 60 * 24 * 2}
	appCtx := component.NewAppContext(db, "Golang", redis, time)

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
	}
	student := r.Group("/student")
	{
		student.POST("/new", studentTransport.AddStudent(appCtx))
		student.GET("/get", studentTransport.GetStudent(appCtx))
		student.POST("/class", studentTransport.AddStudentToClass(appCtx))
		student.POST("/course", studentTransport.AddStudentToCourse(appCtx))
		student.DELETE("/delete", studentTransport.DeleteStudent(appCtx))
		student.DELETE("/delete/class", studentTransport.DeleteStudentFromClass(appCtx))
		student.DELETE("/delete/course", studentTransport.DeleteStudentFromCourse(appCtx))
	}
	result := r.Group("/result")
	{
		result.POST("/new", markTransport.AddResult(appCtx))
		result.PATCH("/update", markTransport.UpdateResult(appCtx))
		result.GET("/list/student", markTransport.ListResultByIdStudent(appCtx))
		result.GET("/list/class", markTransport.ListResultByIdClass(appCtx))
		result.GET("/list/course", markTransport.ListResultByIdCourse(appCtx))
	}

	return r.Run(":8080")
}
