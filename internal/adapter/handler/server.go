package handler

import (
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/gin-gonic/gin"
)

type server struct {
	router  *gin.Engine
	holiday port.HolidayProcessor
}

func NewServer(holiday port.HolidayProcessor) (*server, error) {
	svr := &server{holiday: holiday}
	svr.setupRouter()
	return svr, nil
}

func (svr *server) setupRouter() {
	router := gin.Default()
	router.GET("/api/holidays", svr.holidayHandler)
	svr.router = router
}

func (svr *server) Start(address string) error {
	return svr.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
