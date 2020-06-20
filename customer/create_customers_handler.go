package customer

import (
	"log"
	"net/http"

	"github.com/ckenkub/finalexam/database"
	"github.com/ckenkub/finalexam/errors"
	"github.com/ckenkub/finalexam/types"
	"github.com/gin-gonic/gin"
)

func createCustomerHandler(c *gin.Context) {
	customer := types.Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	customer, err := createCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func createCustomer(customer types.Customer) (types.Customer, error) {
	customer, err := database.CreateCustomer(customer)

	if err != nil {
		log.Println(err)
		return customer, &errors.Error{
			Code:    602,
			Message: "Create data to database error.",
		}
	}

	return customer, err
}
