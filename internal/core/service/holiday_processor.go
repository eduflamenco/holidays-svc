package service

import (
	"context"
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/gookit/slog"
)

type holidayProcessor struct {
	logger *slog.Logger
	data   *[]domain.HoliDay
	filter port.HolidayFilter
}

func NewHolidayProcessor(logger *slog.Logger, data *[]domain.HoliDay, filter port.HolidayFilter) port.HolidayProcessor {
	return &holidayProcessor{logger: logger, data: data, filter: filter}
}

func (p holidayProcessor) ProcessHolidays(ctx context.Context, request domain.HolidayPaymentRequest) (response domain.HoliDayResponse, error error) {

	p.logger.Info("Make payment request", request)

	result, _ := p.filter.Filter(request, *p.data)

	p.logger.Info("Filter result", result)

	return domain.HoliDayResponse{Status: "Success", Holiday: result}, nil
}
