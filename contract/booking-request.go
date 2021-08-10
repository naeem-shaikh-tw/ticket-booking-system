package contract

import "github.com/naeem-shaikh-tw/ticket-booking-system/model"

type BookingRequest struct {
	Catalog model.Catalog `json:"catalog"`
	Slot    model.Slot    `json:"slot"`
}
