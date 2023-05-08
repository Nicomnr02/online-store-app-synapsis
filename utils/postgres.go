package utils

import (
	"online_app_store/model"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:        os.Getenv("DATABASE_URL"),
		DriverName: "pgx",
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(model.Cart{}, model.Category{}, model.Credential{}, model.Product{}, model.Transaction{}, model.User{})
	SetupDBConnection(conn)
	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
