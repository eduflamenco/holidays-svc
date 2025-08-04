package port

import (
	"context"
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
)

type IGetHoliday interface {
	GetHoliday(ctx context.Context) (response []domain.HoliDay, error error)
}

type HolidayProcessor interface {
	ProcessHolidays(ctx context.Context, request domain.HolidayPaymentRequest) (response domain.HoliDayResponse, error error)
}

type HolidayFilter interface {
	Filter(request domain.HolidayPaymentRequest, data []domain.HoliDay) (result []domain.HoliDay, err error)
	SetFilter(filter HolidayFilter)
}
