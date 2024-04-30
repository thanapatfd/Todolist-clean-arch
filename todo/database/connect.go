package database

import (
	"fmt"
	"log"
	"os"

	"github.com/thanapatfd/todolist/todo/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "todo_admin"
	password = "admin"
	dbname   = "postgres"
)

type postgresDB struct {
	Db *gorm.DB
}

func NewPosgrestDB() *postgresDB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&repository.TodoModel{})
	log.Println("running migrations")

	return &postgresDB{Db: db}
}
