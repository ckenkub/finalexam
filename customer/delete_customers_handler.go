package customer

import (
	"log"
	"net/http"

	"github.com/ckenkub/finalexam/database"
	"github.com/ckenkub/finalexam/errors"
	"github.com/gin-gonic/gin"
)

func deleteCustomerHandler(c *gin.Context) {
	id := c.Param("id")

	err := deleteCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}

func deleteCustomerById(id string) error {
	err := database.DeleteCustomerById(id)

	if err != nil {
		log.Println(err)
		return &errors.Error{
			Code:    604,
			Message: "Delete data from database error.",
		}
	}

	return nil
}
