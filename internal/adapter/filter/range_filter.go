package filter

import (
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/gookit/slog"
	"strings"
	"time"
)

type rangeFilter struct {
	filter port.HolidayFilter
	logger *slog.Logger
}

func (t *rangeFilter) SetFilter(filter port.HolidayFilter) {
	t.filter = filter
}

func NewRangeFilter(filter port.HolidayFilter, logger *slog.Logger) port.HolidayFilter {
	return &rangeFilter{filter: filter, logger: logger}
}

func (t *rangeFilter) Filter(request domain.HolidayPaymentRequest, data []domain.HoliDay) (result []domain.HoliDay, err error) {
	filtered := make([]domain.HoliDay, 0)
	if t.filter != nil {
		data, _ = t.filter.Filter(request, data)
	}

	if strings.EqualFold(request.StartDate, "") || strings.EqualFold(request.EndDate, "") {
		return data, nil
	}

	startTime, errstr := time.Parse("2006-01-02", request.StartDate)
	endTime, errend := time.Parse("2006-01-02", request.EndDate)

	if errstr != nil || errend != nil {
		t.logger.Error("Error parsing range dates")
		return filtered, err
	}

	for _, h := range data {
		holidayTime, errH := time.Parse("2006-01-02", h.Date)
		if errH != nil {
			t.logger.Error("Error parsing range date")
			continue
		}
		if holidayTime.Before(startTime) || holidayTime.After(endTime) {
			continue
		}
		filtered = append(filtered, h)
	}
	if len(filtered) == 0 {
		filtered = append(filtered, data...)
	}
	return filtered, err
}
