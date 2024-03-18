package models

type LoginUser struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
