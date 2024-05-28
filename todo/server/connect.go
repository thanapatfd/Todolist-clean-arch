package server

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
	Db *gorm.DB // ตัวแปร Db เก็บการเชื่อมต่อฐานข้อมูล
}


func NewPosgrestDB() *postgresDB {

	// สร้าง Data Source Name (DSN) สำหรับการเชื่อมต่อฐานข้อมูล
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
