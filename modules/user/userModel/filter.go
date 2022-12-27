package userModel

type Filter struct {
	Acp bool `json:"acp,omitempty" bson:"acp"`
	All bool `json:"all,omitempty"`
}
