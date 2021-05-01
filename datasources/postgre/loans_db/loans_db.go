package loans_db

import (
	"log"

	"github.com/bhagyarsh/loan_management/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostDBClient *gorm.DB
	err          error
)

func init() {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=loan port=5432 sslmode=disable"

	PostDBClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	}
	Client, err := PostDBClient.DB()
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database  successfully configured")

	PostDBClient.AutoMigrate(&model.Loan{})

	// var loan = model.Loan{CustomerName: "bhagyarsh", PhoneNo: "8850813167", Email: "Bhayarsh@gmail.com", LoanAmount: 500000, CreditScore: 20}
	// PostDBClient.Create(&loan)

}
