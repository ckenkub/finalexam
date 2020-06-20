package customer

import (
	"log"
	"net/http"

	"github.com/ckenkub/finalexam/database"
	"github.com/ckenkub/finalexam/errors"
	"github.com/ckenkub/finalexam/types"
	"github.com/gin-gonic/gin"
)

func updateCustomerHandler(c *gin.Context) {
	id := c.Param("id")
	customer := types.Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := updateCustomer(customer, id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func updateCustomer(customer types.Customer, id string) error {
	err := database.UpdateCustomer(customer, id)
	if err != nil {
		log.Println(err)
		return &errors.Error{
			Code:    603,
			Message: "Update data to database error.",
		}
	}

	return nil
}
