package customer

import (
	"log"
	"net/http"

	"github.com/ckenkub/finalexam/database"
	"github.com/ckenkub/finalexam/errors"
	"github.com/ckenkub/finalexam/types"
	"github.com/gin-gonic/gin"
)

func getCustomersHandler(c *gin.Context) {
	customers, err := getCustomers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomerHandler(c *gin.Context) {
	id := c.Param("id")

	customer, err := getCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func getCustomers() ([]*types.Customer, error) {
	customers, err := database.GetCustomers()

	if err != nil {
		log.Println(err)
		return customers, &errors.Error{
			Code:    601,
			Message: "Get data from database error.",
		}
	}

	return customers, nil
}

func getCustomerById(id string) (types.Customer, error) {
	customers, err := database.GetCustomerByID(id)

	if err != nil {
		log.Println(err)
		return customers, &errors.Error{
			Code:    601,
			Message: "Get data from database error.",
		}
	}

	return customers, nil
}
