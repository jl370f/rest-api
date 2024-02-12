package config

import (
	"fmt"
	"golang-api/apiError"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", viper.GetString("app.host"), viper.GetInt("app.port"), viper.GetString("app.user"), viper.GetString("app.password"), viper.GetString("app.dbName"))
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	apiError.HandleError(err)

	return db
}
