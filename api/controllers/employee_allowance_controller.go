package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/t2469/labor-management-system.git/db"
	"github.com/t2469/labor-management-system.git/helpers"
	"github.com/t2469/labor-management-system.git/models"
	"net/http"
)

func CreateEmployeeAllowance(c *gin.Context) {
	var ea models.EmployeeAllowance
	if err := c.ShouldBindJSON(&ea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID, err := helpers.GetCompanyID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var emp models.Employee
	if err := db.DB.First(&emp, ea.EmployeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if emp.CompanyID != companyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed to create allowance for this employee"})
		return
	}

	if err := db.DB.Create(&ea).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ea)
}

func GetEmployeeAllowances(c *gin.Context) {
	companyID, err := helpers.GetCompanyID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var employees []models.Employee
	if err := db.DB.Where("company_id = ?", companyID).Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var allAllowances []models.EmployeeAllowance
	for _, emp := range employees {
		var allowances []models.EmployeeAllowance
		if err := db.DB.Where("employee_id = ?", emp.ID).Find(&allowances).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allAllowances = append(allAllowances, allowances...)
	}

	c.JSON(http.StatusOK, allAllowances)
}
