package main

import (
	"context"
	"github.com/eduflamenco/holidays-svc/internal/adapter/client/rest"
	"github.com/eduflamenco/holidays-svc/internal/adapter/filter"
	"github.com/eduflamenco/holidays-svc/internal/adapter/handler"
	"github.com/eduflamenco/holidays-svc/internal/core/service"
	"github.com/eduflamenco/holidays-svc/internal/core/util/client"
	"github.com/eduflamenco/holidays-svc/internal/core/util/config"
	"github.com/gookit/slog"
	shandler "github.com/gookit/slog/handler"
	"log"
	"time"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations: ", err.Error())
	}
	logger := NewCustomLogger()
	restClient := client.NewHttpClient(1 * time.Minute)
	getHoliday := rest.NewGetHolidays(logger, config, client.NewRestClient(restClient))
	data, err := getHoliday.GetHoliday(context.Background())
	if err != nil {
		log.Fatal("Cannot get holidays: ", err.Error())
	}
	filters := filter.NewFilterFactory(logger, config).CreateFilterDecorator()
	retriever := service.NewHolidayProcessor(logger, &data, filters)
	server, err := handler.NewServer(retriever)
	if err != nil {
		log.Fatal(err)
	}
	server.Start(config.HTTPServerAddress)

}

func NewCustomLogger() *slog.Logger {
	l := slog.New()
	l.Config()
	// add handlers ...
	h1 := shandler.NewConsoleHandler(slog.AllLevels)
	h1.SetFormatter(slog.NewJSONFormatter())
	l.AddHandlers(h1)
	return l
}
