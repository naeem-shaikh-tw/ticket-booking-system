package contract

import "github.com/naeem-shaikh-tw/ticket-booking-system/model"

type BookingResponse struct {
	Success bool           `json:"success"`
	Errors  []string       `json:"errors"`
	Data    []model.Ticket `json:"data"`
}
