package handler

import (
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) holidayHandler(ctx *gin.Context) {

	holidayType := ctx.Query("type")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	req := domain.HolidayPaymentRequest{HolidayType: holidayType, StartDate: startDate, EndDate: endDate}
	response, err := s.holiday.ProcessHolidays(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if ctx.GetHeader("Content-Type") == "application/xml" {
		ctx.XML(http.StatusOK, response)
	}
	ctx.JSON(http.StatusOK, response)

}
