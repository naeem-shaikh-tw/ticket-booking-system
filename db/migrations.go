package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/naeem-shaikh-tw/ticket-booking-system/db/seed"
)

func RunMigrations() {
	fmt.Println("Running migrations")

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error in migration", err))
	}
	fmt.Println("got the driver")
	m, error := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres", driver)

	if error != nil {
		panic(error)
	}
	m.Up()
	if err != nil {
		panic(fmt.Sprintf("Error : %+v \n", err))
	}
	fmt.Println("Migration SUCCESS")
}

func RollbackMigrations() {
	fmt.Println("Running Rollback migrations")

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error in migration", err))
	}
	fmt.Println("got the driver")
	m, error := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres", driver)

	if error != nil {
		panic(error)
	}
	m.Down()
	if err != nil {
		panic(fmt.Sprintf("Error : %+v \n", err))
	}
	fmt.Println("Rollback SUCCESS")
}

func Seed() {
	for _, query := range seed.MIGRATIONS {
		_, err := DbPool.Exec(query)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Seed migration Done")
}
