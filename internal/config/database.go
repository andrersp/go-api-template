package config

import (
	"log"

	"github.com/andrersp/go-api-template/internal/domain/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateSQLiteConn() (err error) {

	if dbConn != nil {
		CloseConnection(dbConn)
	}

	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbConn = db
	return
}

// func CreatePostgresConn() (err error) {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
// 		DB_HOST, DB_USER, DB_PASSWD, DB_NAME, DB_PORT, DB_SSL_MODE)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})

// 	if err != nil {
// 		return
// 	}

// 	dbConn = db
// 	return
// }

func ConnectDB() (db *gorm.DB, err error) {

	sqlDB, err := dbConn.DB()

	if sqlDB == nil {
		return
	}

	if err = sqlDB.Ping(); err != nil {
		return
	}

	db = dbConn
	return
}

func CloseConnection(conn *gorm.DB) {
	db, err := conn.DB()
	if err != nil {
		return
	}

	defer db.Close()
}

func AutoMigrate() error {
	db, err := ConnectDB()

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&user.User{})

	return err
}
