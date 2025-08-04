package rest

import (
	"context"
	"fmt"
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/eduflamenco/holidays-svc/internal/core/port"
	"github.com/eduflamenco/holidays-svc/internal/core/util/client"
	"github.com/eduflamenco/holidays-svc/internal/core/util/config"
	"github.com/gookit/slog"
	"strings"
)

type getHolidays struct {
	logger     *slog.Logger
	conf       config.Config
	restClient *client.RestClient
}

func NewGetHolidays(logger *slog.Logger, conf config.Config, restClient *client.RestClient) port.IGetHoliday {
	return &getHolidays{logger: logger, conf: conf, restClient: restClient}
}

func (s *getHolidays) GetHoliday(ctx context.Context) (response []domain.HoliDay, error error) {
	var res domain.HoliDayResponse
	s.logger.Info("Get Holidays request", s.conf.FeriadosApiUrl)
	err := s.restClient.GetRequest(ctx, s.conf.FeriadosApiUrl, make(map[string]string), &res)
	if err != nil {
		s.logger.Errorf("Error getting holidays info: %v", err)
		return res.Holiday, err
	}
	if !strings.EqualFold(res.Status, "success") {
		s.logger.Error("Holiday status is not success")
		return res.Holiday, fmt.Errorf("holiday status is not success")
	}
	s.logger.Info("Getting Holidays response", res)
	return res.Holiday, nil
}
