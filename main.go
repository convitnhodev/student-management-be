package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"managerstudent/common/setupDatabase"
	"managerstudent/component"
	"managerstudent/modules/user/userTransport"
)

func main() {
	db := setupDatabase.InitMongoDB()
	fmt.Println(db)
	if err := runService(db, nil); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *mongo.Client, redis *redis.Client) error {
	r := gin.Default()

	time := component.TimeJWT{60 * 60 * 24 * 2, 60 * 60 * 24 * 2}
	appCtx := component.NewAppContext(db, "anhHaudungboemnhe", redis, time)

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))

	}

	return r.Run(":8080")
}
