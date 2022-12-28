package userModel

type Filter struct {
	Acp bool `json:"acp,omitempty" bson:"acp" header:"acp"`
	All bool `json:"all,omitempty" header:"all"`
}
