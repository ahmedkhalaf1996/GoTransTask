package handlers

import (
	"net/http"
	"time"

	"github.com/freddymac124/upwork-go/models"
	"github.com/gofiber/fiber/v2"
)

type AddDebit struct {
	Description string  `json:"description" binding:"required"`
	Debit       float64 `json:"debit" binding:"required"`
}

type AddCredit struct {
	Description string  `json:"description" binding:"required"`
	Credit      float64 `json:"credit" binding:"required"`
}

func AddNewDebit(c *fiber.Ctx) error {

	request := &AddDebit{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// first, check if the user is exist
	transaction := models.Transaction{}
	models.DB.Last(&transaction)

	updated := models.Transaction{
		Description: request.Description,
		Debit:       int(request.Debit) - transaction.Debit,
		Credit:      0,
		Balance:     transaction.Balance - int(request.Debit),
		Date:        time.Now(),
	}
	models.DB.Create(&updated)

	var transactionRecored []models.Transaction
	models.DB.Order("id DESC").Limit(5).Find(&transactionRecored)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"transactionRecored": transactionRecored,
	})

}

func AddNewCredit(c *fiber.Ctx) error {

	request := &AddCredit{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	// first, check if the user is exist
	transaction := models.Transaction{}
	models.DB.Last(&transaction)

	updated := models.Transaction{
		Description: request.Description,
		Credit:      int(request.Credit) + transaction.Credit,
		Debit:       0,
		Balance:     transaction.Balance + int(request.Credit),
		Date:        time.Now(),
	}
	models.DB.Create(&updated)

	var transactionRecored []models.Transaction
	models.DB.Order("id DESC").Limit(5).Find(&transactionRecored)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"transactionRecored": transactionRecored,
	})

}

func GetLastRecored(c *fiber.Ctx) error {
	var transactionRecored []models.Transaction
	models.DB.Order("id DESC").Limit(5).Find(&transactionRecored)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"transactionRecored": transactionRecored,
	})
}
