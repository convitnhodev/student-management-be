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
	"managerstudent/middleware"
	"managerstudent/modules/class/classTransport"
	"managerstudent/modules/course/courseTransport"
	"managerstudent/modules/result/resultTransport"
	"managerstudent/modules/student/studentTransport"
	"managerstudent/modules/subject/subjectTransport"
	"managerstudent/modules/user/userTransport"
	"managerstudent/rules"
)

func main() {
	managerLog.InitLogs()
	db := setupDatabase.InitMongoDB()
	redis := setupDatabase.InitRedis()
	fmt.Println(db)
	if err := runService(db, redis); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *mongo.Client, redis *redis.Client) error {
	r := gin.Default()
	time := component.TimeJWT{60 * 60 * 24 * 2, 60 * 60 * 24 * 2}
	appCtx := component.NewAppContext(db, "Golang", redis, time, localPubsub.NewPubSub())
	r.Use(middleware.CORSMiddleware())

	subjectRoute := r.Group("/subject")
	{
		subjectRoute.POST("/new", subjectTransport.NewCreateSubject(appCtx))
		subjectRoute.DELETE("/delete", subjectTransport.DeleteSubject(appCtx))
		subjectRoute.GET("/list", subjectTransport.ListSubjects(appCtx))
	}

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
		user.POST("/login", userTransport.Login(appCtx))
		user.GET("/list", userTransport.ListUsers(appCtx))
		user.GET("/get", userTransport.GetByUsername(appCtx))
		user.POST("/update/homeroom", userTransport.UpdateHomeroom(appCtx))
		user.POST("/update/teaching", userTransport.UpdateTeaching(appCtx))
		user.PATCH("/update", userTransport.UpdateUser(appCtx))
		user.PATCH("/accept", userTransport.AcceptUser(appCtx))
		user.PATCH("/update/password", userTransport.UserUpdatePassword(appCtx))
		user.GET("/profile", middleware.RequireAuth(appCtx), userTransport.GetProfile(appCtx))
	}
	student := r.Group("/student")
	{
		student.POST("/new", studentTransport.AddStudent(appCtx))
		student.PATCH("/update", studentTransport.UpdateStudent(appCtx))
		student.GET("/get", studentTransport.GetStudent(appCtx))
		//student.POST("/class", studentTransport.AddStudentToClass(appCtx))
		//student.POST("/course", studentTransport.AddStudentToCourse(appCtx))
		//student.DELETE("/delete", studentTransport.DeleteStudent(appCtx))
		//student.DELETE("/delete/class", studentTransport.DeleteStudentFromClass(appCtx))
		//student.DELETE("/delete/course", studentTransport.DeleteStudentFromCourse(appCtx))
	}
	result := r.Group("/result")
	{
		result.POST("/new", resultTransport.AddResult(appCtx))
		result.DELETE("/delete", resultTransport.DeleteResults(appCtx))
		result.PATCH("/update", resultTransport.UpdateResult(appCtx))
	}

	admin := r.Group("/admin", middleware.RequireAuth(appCtx), middleware.RequireRole(appCtx, 1))
	{
		rulesRoute := admin.Group("/rules")
		{
			rulesRoute.GET("/get", rules.GetRules(appCtx))
			rulesRoute.POST("/update", rules.UpdateRules(appCtx))
		}

		course := admin.Group("course")
		{
			course.POST("/new", courseTransport.CreateNewCourse(appCtx))
			course.DELETE("/delete", courseTransport.DeleteCourse(appCtx))
			course.GET("/list", courseTransport.ListCourses(appCtx))
		}


		class := admin.Group("/class")
		{
			class.POST("/new", classTransport.CreateNewClass(appCtx))
			class.DELETE("/delete", classTransport.DeleteClass(appCtx))
			class.GET("/list", classTransport.ListClasses(appCtx))
			class.GET("/list/student", classTransport.ListStudents(appCtx))
			class.GET("/get", classTransport.GetClass(appCtx))
		}
	}


	return r.Run(":8080")
}
