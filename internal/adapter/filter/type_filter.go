package filter

import (
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/gookit/slog"
	"strings"
)

type typeFilter struct {
	filter port.HolidayFilter
	logger *slog.Logger
}

func (t *typeFilter) SetFilter(filter port.HolidayFilter) {
	t.filter = filter
}

func NewTypeFilter(filter port.HolidayFilter, logger *slog.Logger) port.HolidayFilter {
	return &typeFilter{filter: filter, logger: logger}
}

func (t *typeFilter) Filter(request domain.HolidayPaymentRequest, data []domain.HoliDay) (result []domain.HoliDay, err error) {
	filtered := make([]domain.HoliDay, 0)
	if t.filter != nil {
		data, _ = t.filter.Filter(request, data)
	}

	for _, h := range data {
		if strings.EqualFold(request.HolidayType, h.Type) {
			filtered = append(filtered, h)
		}
	}
	if len(filtered) == 0 {
		filtered = append(filtered, data...)
	}
	return filtered, err
}
