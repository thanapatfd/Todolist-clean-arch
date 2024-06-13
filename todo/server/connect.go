package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thanapatfd/todolist/todo/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDB struct {
	Db *gorm.DB // ตัวแปร Db เก็บการเชื่อมต่อฐานข้อมูล
}

func NewPosgrestDB() *postgresDB {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	dbHostname := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	pgAdminEmail := os.Getenv("PGADMIN_DEFAULT_EMAIL")
	pgAdminPassword := os.Getenv("PGADMIN_DEFAULT_PASSWORD")
	port := os.Getenv("PORT")

	fmt.Println("DB Hostname:", dbHostname)
	fmt.Println("Database Name:", dbName)
	fmt.Println("Database User:", dbUser)
	fmt.Println("Database Password:", dbPassword)
	fmt.Println("PgAdmin Email:", pgAdminEmail)
	fmt.Println("PgAdmin Password:", pgAdminPassword)

	// สร้าง Data Source Name (DSN) สำหรับการเชื่อมต่อฐานข้อมูล
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHostname, port, dbUser, dbPassword, dbName)

	// เปิดการเชื่อมต่อกับฐานข้อมูล Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // ตั้งค่า logger ให้แสดงข้อมูลระดับ Info
	})

	// ตรวจสอบว่าการเชื่อมต่อฐานข้อมูลสำเร็จหรือไม่
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&repository.TodoModel{}) // รันการทำ migrations สำหรับ model TodoModel
	log.Println("running migrations")

	return &postgresDB{Db: db} // คืนค่า instance ของ postgresDB ที่มีการเชื่อมต่อฐานข้อมูล
}
