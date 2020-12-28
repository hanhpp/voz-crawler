package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type VecConfig struct {
	ServerPort               string
	ServerMode               string
	MaxLimit                 int
	PostgresConnectionString string
	RedisConnectionHost      string
	RedisConnectionPassword  string
}

var Config *VecConfig

func FetchEnvironmentVariables() {
	Config = NewVecConfig()
	Config.GetConfig()
}

func NewVecConfig() *VecConfig {
	cf := VecConfig{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return &cf
}

func (config *VecConfig) GetConfig() {
	port := os.Getenv("PORT")
	config.ServerPort = port
	ginMode := strings.TrimSpace(strings.ToLower(os.Getenv("GIN_ENV")))
	switch ginMode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		gin.SetMode(ginMode)
		config.ServerMode = ginMode
	default:
		gin.SetMode(gin.DebugMode)
		config.ServerMode = gin.DebugMode
	}

	var dbHost = os.Getenv("DATABASE_HOST")
	var dbPort = os.Getenv("DATABASE_PORT")
	var username = os.Getenv("DATABASE_USERNAME")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbName = os.Getenv("DATABASE_NAME")
	var sslMode = os.Getenv("DATABASE_SSL_MODE")

	config.PostgresConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, username, dbName, password, sslMode)
	config.MaxLimit, _ = strconv.Atoi(os.Getenv("MAX_LIMIT"))

	config.RedisConnectionHost = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	config.RedisConnectionPassword = os.Getenv("REDIS_PASSWORD")
}
