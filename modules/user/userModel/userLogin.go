package userModel

type UserLogin struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
