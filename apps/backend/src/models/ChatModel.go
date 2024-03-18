package models

type Message struct {
	Username string `json:"username" bson:"username"`
	Message  string `json:"message" bson:"message"`
	Type    string `json:"type" bson:"type"`
   }