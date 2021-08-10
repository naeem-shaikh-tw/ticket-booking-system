package app

import (
	"github.com/gin-gonic/gin"
	"github.com/naeem-shaikh-tw/ticket-booking-system/handler"
)

type Router struct{ engine *gin.Engine }

func CreateRoutes() *gin.Engine {

	router := Router{engine: gin.Default()}
	router.registerRoutes()
	return router.engine
}

func (router Router) registerRoutes() {
	router.engine.GET("/", handler.PingHandler)
	router.engine.GET("/book", handler.GetTicketHandler)
	router.engine.POST("/book", handler.PostBookHandler)
}

// type Router struct{ *gin.Engine }

// func CreateRoutes() Router {

// 	router := &Router{gin.Default()}
// 	router.registerRoutes()
// 	return *router
// }

// func (router Router) registerRoutes() {
// 	router.GET("/", handler.PingHandler)
// 	router.POST("/book", handler.PostBookHandler)
// }
