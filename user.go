package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "ez_api"
	password = "somepwd123"
	dbname   = "ez_api_db"
)

var DB *gorm.DB
var err error

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
var DNS = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

type Student struct {
	gorm.Model
	Name        string `json: "name"`
	Age         int    `json: age`
	Grade       string `json: grade`
	Description string `json: description`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic("Cannot connect to DB")
	}

	DB.AutoMigrate(&Student{})
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []Student
	DB.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	DB.Create(&student)
	json.NewEncoder(w).Encode(student)
}
