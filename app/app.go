package app

import (
	"fmt"
)

func StartServer() {
	router := CreateRoutes()

	fmt.Println("Registered Routes, STarting the server")
	router.Run()
}
