package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naeem-shaikh-tw/ticket-booking-system/contract"
	"github.com/naeem-shaikh-tw/ticket-booking-system/model"
)

func PostBookHandler(c *gin.Context) {

	var requestBody contract.BookingRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		panic(fmt.Sprint("error - ", err))
	}
	ticket := model.CreateTicket(requestBody.Catalog, requestBody.Slot)

	c.JSON(http.StatusCreated, contract.BookingResponse{
		Data: []model.Ticket{ticket},
	})
}

func GetTicketHandler(c *gin.Context) {

	ticketIdParam := c.Request.URL.Query().Get("ticketId")
	ticketId, err := strconv.ParseInt(ticketIdParam, 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	ticket, err := model.LoadTicket(int(ticketId))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, contract.BookingResponse{
		Data: []model.Ticket{ticket},
	})
}
