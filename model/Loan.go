package model

import (
	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	CustomerName string  `json:"name"`
	PhoneNo      string  `json:"phoneNo"`
	Email        string  `json:"email"`
	LoanAmount   float64 `json:"loanAmount"`
	CreditScore  float32 `json:"creditScore"`
	LoanStatus   string  `json:"status"`
}

type Status struct {
	Status string `json:"status"`
}
