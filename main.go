package main

import (
	"fmt"
	"log"
	"managerstudent/common/pubsub/localPubsub"
	"managerstudent/common/setupDatabase"
	"managerstudent/component"
	"managerstudent/component/managerLog"
	"managerstudent/middleware"
	"managerstudent/modules/class/classTransport"
	"managerstudent/modules/course/courseTransport"
	"managerstudent/modules/notifedProvider/notificationTransport"
	"managerstudent/modules/result/resultTransport"
	"managerstudent/modules/student/studentTransport"
	"managerstudent/modules/subcriber"
	"managerstudent/modules/user/userTransport"
	"managerstudent/rules"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
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
	r.Use(middleware.CORSMiddleware())
	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
		user.POST("/login", userTransport.Login(appCtx))
		user.GET("/list", userTransport.ListUsers(appCtx))
		user.GET("/get", userTransport.GetByUsername(appCtx))
		user.PATCH("/update", userTransport.UpdateUser(appCtx))
		user.PATCH("/accept", userTransport.AcceptUser(appCtx))
		user.PATCH("/update/password", userTransport.UserUpdatePassword(appCtx))
	}
	student := r.Group("/student")
	{
		student.POST("/new", studentTransport.AddStudent(appCtx))
		student.POST("/add/class", studentTransport.UserAddStudentToClass(appCtx))
		student.POST("/add/course", studentTransport.UserAddStudentToCourse(appCtx))
		//student.PATCH("/update/result", studentTransport.UserUpdateResult(appCtx))
		//student.GET("/get", studentTransport.GetStudent(appCtx))
		//student.POST("/class", studentTransport.AddStudentToClass(appCtx))
		//student.POST("/course", studentTransport.AddStudentToCourse(appCtx))
		//student.DELETE("/delete", studentTransport.DeleteStudent(appCtx))
		//student.DELETE("/delete/class", studentTransport.DeleteStudentFromClass(appCtx))
		//student.DELETE("/delete/course", studentTransport.DeleteStudentFromCourse(appCtx))
	}
	result := r.Group("/result")
	{
		//result.POST("/new", resultTransport.AddResult(appCtx))
		result.POST("/new", resultTransport.AddResult(appCtx))
		result.DELETE("/delete", resultTransport.DeleteResults(appCtx))
		result.PATCH("/update", resultTransport.UpdateResult(appCtx))
		//result.GET("/list/student", resultTransport.ListResultByIdStudent(appCtx))
		//result.GET("/list/class", resultTransport.ListResultByIdClass(appCtx))
		//result.GET("/list/course", resultTransport.ListResultByIdCourse(appCtx))
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

	notify := r.Group("/notification")
	{
		notify.GET("/get", notificationTransport.GetNotification(appCtx))
		notify.GET("/list", notificationTransport.ListNotifications(appCtx))
		notify.POST("/acp/user", notificationTransport.AdminAcpNotifyUserRegister(appCtx))
		notify.POST("/acp/student", notificationTransport.AdminAcpNotifyRequestAddStudent(appCtx))
	}

	rulesRoute := r.Group("/rules")
	{
		rulesRoute.GET("/get", rules.GetRules(appCtx))
		rulesRoute.POST("/update", rules.UpdateRules(appCtx))
	}
	return r.Run(":8080")
}
