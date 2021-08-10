package main

import (
	"flag"

	"github.com/naeem-shaikh-tw/ticket-booking-system/app"
	"github.com/naeem-shaikh-tw/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	rollback := flag.Bool("rollback", false, "To run rollback")
	seed := flag.Bool("seed", false, "To seed data")
	flag.Parse()

	dbConfig := db.Db{

		Name:     "ticketbookingdb",
		Host:     "localhost",
		Port:     5432,
		UserName: "ticketbookinguser",
		Password: "AToughPassword!",
	}
	db.DbPool = dbConfig.Connect()

	if *migrate {
		db.RunMigrations()
		return
	}
	if *rollback {
		db.RollbackMigrations()
		return
	}

	if *seed {
		db.Seed()
		return
	}

	app.StartServer()
}
