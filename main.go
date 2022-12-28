package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"managerstudent/common/pubsub/localPubsub"
	"managerstudent/common/setupDatabase"
	"managerstudent/component"
	"managerstudent/component/managerLog"
	"managerstudent/modules/class/classTransport"
	"managerstudent/modules/course/courseTransport"
	"managerstudent/modules/mark/markTransport"
	"managerstudent/modules/notifedProvider/notifyTransport"
	"managerstudent/modules/student/studentTransport"
	"managerstudent/modules/subcriber"
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
	appCtx := component.NewAppContext(db, "Golang", redis, time, localPubsub.NewPubSub())
	subcriber.Setup(appCtx)
	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
		user.POST("/login", userTransport.Login(appCtx))
		user.GET("list", userTransport.ListUsers(appCtx))
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

	class := r.Group("/class")
	{
		class.POST("/new", classTransport.CreateNewClass(appCtx))
		class.DELETE("/delete", classTransport.DeleteClass(appCtx))
		class.GET("/list", classTransport.ListClasses(appCtx))
	}

	course := r.Group("course")
	{
		course.POST("/new", courseTransport.CreateNewCourse(appCtx))
		course.DELETE("/delete", courseTransport.DeleteCourse(appCtx))
		course.GET("/list", courseTransport.ListCourses(appCtx))
	}

	notify := r.Group("/notify")
	{
		notify.GET("/get", notifyTransport.GetNotification(appCtx))
		notify.GET("/list", notifyTransport.ListNotifications(appCtx))
		notify.POST("/acp/user", notifyTransport.AdminAcpNotifyUserRegister(appCtx))
		notify.POST("/acp/student", notifyTransport.AdminAcpNotifyRequestAddStudent(appCtx))
	}

	return r.Run(":8080")
}
