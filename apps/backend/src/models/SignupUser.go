package models

type SignupUser struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
}
