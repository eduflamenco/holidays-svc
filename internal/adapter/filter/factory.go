package filter

import (
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/eduflamenco/holidays-svc/internal/core/util/config"
	"github.com/gookit/slog"
	"strings"
)

const (
	Type  = "type"
	Range = "range"
)

type filterFactory struct {
	conf   config.Config
	logger *slog.Logger
}

func NewFilterFactory(logger *slog.Logger, conf config.Config) *filterFactory {
	return &filterFactory{logger: logger, conf: conf}
}

func (f *filterFactory) CreateFilterDecorator() port.HolidayFilter {
	var currentImpl port.HolidayFilter
	var finalImpl port.HolidayFilter
	var exists bool
	filters := f.getExistingFilterDecorators()
	fltOrder := strings.Split(f.conf.FilterOrder, ",")
	for _, v := range fltOrder {
		if currentImpl, exists = filters[strings.TrimSpace(v)]; !exists {
			f.logger.Error("Method: CreateViewsDecorators", "No existe un filtro con el nombre", v)
			continue
		}
		currentImpl.SetFilter(finalImpl)
		finalImpl = currentImpl
	}
	return finalImpl
}

func (f *filterFactory) getExistingFilterDecorators() map[string]port.HolidayFilter {

	filters := make(map[string]port.HolidayFilter)
	filters[Type] = NewTypeFilter(nil, f.logger)
	filters[Range] = NewRangeFilter(nil, f.logger)
	return filters
}
