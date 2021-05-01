package services

import (
	"log"
	"strings"

	"github.com/bhagyarsh/loan_management/datasources/postgre/loans_db"
	"github.com/bhagyarsh/loan_management/model"
)

func GetAlloanFromDB(status string, loanAmountGreater string, records *[]model.Loan) {
	if status != "" && loanAmountGreater == "" {
		res1 := strings.Split(status, ",")
		loans_db.PostDBClient.Where("loan_status IN ?", res1).Find(&records)

		return
	} else if status == "" && loanAmountGreater != "" {
		loans_db.PostDBClient.Where("loan_amount > ?", loanAmountGreater).Find(&records)
		return
	} else if status != "" && loanAmountGreater != "" {
		res1 := strings.Split(status, ",")
		loans_db.PostDBClient.Where("loan_amount > ? AND loan_status IN ?", loanAmountGreater, res1).Find(&records)

		return
	} else {
		log.Println(status)

		if err := loans_db.PostDBClient.Find(&records).Error; err != nil {
			log.Fatalln(err)
		}

	}
}

func GetloanByID(id string, loan *model.Loan) error {
	err := loans_db.PostDBClient.First(&loan, id).Error
	if err != nil {
		return err
	}
	log.Println(loan.CustomerName)

	return nil
}

func ApproveloanByID(id string, ls *model.Status) error {
	if (ls.Status == "new") || (ls.Status == "approved") || (ls.Status == "cancelled") || (ls.Status == "rejected") {
		err := loans_db.PostDBClient.First(&model.Loan{}, id).Error
		if err != nil {

			return err
		}
		err = loans_db.PostDBClient.Model(&model.Loan{}).Where("ID = ?", id).Update("loan_status", ls.Status).Error
		if err != nil {

			return err
		}
	}
	return nil
}
