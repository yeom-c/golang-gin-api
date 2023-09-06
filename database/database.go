package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/yeom-c/golang-gin-api/config"
	"github.com/yeom-c/golang-gin-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB struct {
	Gorm   *gorm.DB
	Dynamo *dynamodb.Client
}

func init() {
	initGorm()
	initDynamoDB()
}

func initGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.EnvVar.DBUser, cfg.EnvVar.DBPassword, cfg.EnvVar.DBHost, cfg.EnvVar.DBPort, cfg.EnvVar.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// auto migration.
	db.AutoMigrate(&model.Account{})

	DB.Gorm = db
}

func initDynamoDB() {
	awsConfig, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cfg.EnvVar.DynamoDBRegion), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.EnvVar.DynamoDBAccessKeyId, cfg.EnvVar.DynamoDBSecretAccessKey, "")))
	if err != nil {
		log.Fatal("failed to load aws config: ", err)
	}

	client := dynamodb.NewFromConfig(awsConfig)
	DB.Dynamo = client
}
