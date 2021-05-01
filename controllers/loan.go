package controllers

import (
	"log"
	"net/http"

	"github.com/bhagyarsh/loan_management/datasources/postgre/loans_db"
	"github.com/bhagyarsh/loan_management/model"
	"github.com/bhagyarsh/loan_management/services"
	"github.com/gin-gonic/gin"
)

func GetAlloan(c *gin.Context) {
	var records []model.Loan
	status := c.Query("status")
	loanAmountGreater := c.Query("loanAmountGreater")

	services.GetAlloanFromDB(status, loanAmountGreater, &records)
	c.JSON(http.StatusOK, gin.H{
		"result": &records,
	})

}
func GetloanByID(c *gin.Context) {
	loanid := c.Param("id")
	log.Println(loanid)
	var loan model.Loan
	err := services.GetloanByID(loanid, &loan)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "loan with id not found", // cast it to string before showing
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": &loan, // cast it to string before showing
	})
}

func ApproveloanByID(c *gin.Context) {
	loanid := c.Param("id")
	var loanStatus model.Status
	err := c.BindJSON(&loanStatus)
	if err != nil {

		return
	}
	err = services.ApproveloanByID(loanid, &loanStatus)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "loan with id not found", // cast it to string before showing
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "changed status of loan",
	})

}

func Createloan(c *gin.Context) {

	var loan model.Loan
	err := c.BindJSON(&loan)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please check data",
		})
		return
	}

	result := loans_db.PostDBClient.Create(&loan)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "unable to create loan",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"loan": &loan,
	})
}

func Cancelloan(c *gin.Context) {
	loanID := c.Param("id")

	err := loans_db.PostDBClient.Delete(&model.Loan{}, loanID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "loan id not found", // cast it to string before showing
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "loan canceled", // cast it to string before showing
	})
}
