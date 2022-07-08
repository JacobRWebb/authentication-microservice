package database

import (
	"fmt"

	"github.com/JacobRWebb/authentication-microservice/database/models"
	"github.com/JacobRWebb/authentication-microservice/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseConnection() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", util.GetEnvVariable("POSTGRES_HOST"), util.GetEnvVariable("POSTGRES_USER"), util.GetEnvVariable("POSTGRES_PASSWORD"), util.GetEnvVariable("POSTGRES_DB"), util.GetEnvVariable("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
	}

	db.AutoMigrate(&models.User{})

	return db;
}