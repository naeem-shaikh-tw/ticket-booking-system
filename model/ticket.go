package model

import (
	"log"

	"github.com/naeem-shaikh-tw/ticket-booking-system/db"
)

type Ticket struct {
	Id        int `json:"id"`
	CatalogId int `json:"catalog"`
	SlotId    int `json:"slot"`
}

var querySet = db.QuerySet{
	SelectQuery: "SELECT slot_id, catalog_id FROM tickets WHERE id = $1 ;",
	InsertQuery: "INSERT INTO tickets ( catalog_id, slot_id ) VALUES ( $1, $2 ) RETURNING id;",
}

func CreateTicket(c Catalog, s Slot) Ticket {

	stmt, err := db.DbPool.Prepare(querySet.InsertQuery)
	if err != nil {
		log.Fatal("Error while preparing statement")
	}
	defer stmt.Close()

	var ticketId int
	err = stmt.QueryRow(c.Id, s.Id).Scan(&ticketId)
	if err != nil {
		log.Fatal(err)
	}
	return Ticket{
		SlotId:    s.Id,
		CatalogId: c.Id,
		Id:        ticketId,
	}
}

func LoadTicket(ticketId int) (ticket Ticket, error error) {
	stmt, err := db.DbPool.Prepare(querySet.SelectQuery)
	if err != nil {
		log.Fatal("Error while preparing statement")
	}
	defer stmt.Close()

	var slotId int
	var catalogId int
	err = stmt.QueryRow(ticketId).Scan(&slotId, &catalogId)
	if err != nil {
		log.Fatal(err)
	}
	ticket = Ticket{
		SlotId:    slotId,
		CatalogId: catalogId,
		Id:        ticketId,
	}
	return
}
