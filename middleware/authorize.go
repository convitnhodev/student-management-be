package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	_const "managerstudent/common/const"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/component/tokenProvider/jwt"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"
	"strings"
)

func ErrWrongAuthHeader(err error) *solveError.AppError {
	return solveError.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	// split Bearer and token
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecret())

	return func(c *gin.Context) {

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}



		redisClient := appCtx.GetRedis()

		data,err := redisClient.Get(token).Bytes()
		var resultData userModel.User
		err = json.Unmarshal(data, &resultData)
		if err != nil {
			panic(err)
		}

		if err == nil {
			c.Set(_const.CurrentUser, &resultData)
			c.Next()
			return
		}

		db := appCtx.GetNewDataMongoDB()
		store := userStorage.NewMongoStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.JSON(401, "token is invalid")
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), bson.M{"username": payload.UserName})
		if err != nil {
			panic(err)
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		err = redisClient.Set(token, jsonUser, 3600 * 24 * 3).Err()
		if err != nil {
			panic(err)
		}

		c.Set(_const.CurrentUser, user)
		c.Next()
	}
}
