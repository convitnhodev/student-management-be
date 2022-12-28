package userModel

type UserLogin struct {
	UserName string `json:"user_name" bson:"user_name"`
	Password string `json:"pass_word" bson:"pass_word"`
}
