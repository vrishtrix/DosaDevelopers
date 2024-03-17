package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var (
	PORT         string = os.Getenv("PORT")
	MONGO_URL    string = os.Getenv("MONGO_URL")
	SECRET_KEY   string = os.Getenv("SECRET_KEY")
	DEFAULT_ROLE string = os.Getenv("DEFAULT_ROLE")
	COUNTER_USER int64
	COUNTER_POST int64
	MU           sync.Mutex
)
