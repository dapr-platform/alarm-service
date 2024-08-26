package main

import (
	"alarm-service/api"
	"alarm-service/collect"
	_ "alarm-service/docs"
	"alarm-service/event"
	"github.com/dapr-platform/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"strconv"
)

var (
	PORT = 80
)

func init() {

	if val := os.Getenv("LISTEN_PORT"); val != "" {
		PORT, _ = strconv.Atoi(val)
	}
	common.Logger.Debug("use PORT ", PORT)
}

// @title alarm-service API
// @version 1.0
// @description alarm-service API
// @BasePath /swagger/alarm-service
func main() {
	mux := chi.NewRouter()
	api.InitRoute(mux)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/swagger*", httpSwagger.WrapHandler)
	s := daprd.NewServiceWithMux(":"+strconv.Itoa(PORT), mux)
	event.NewEventHandler(s, common.PUBSUB_NAME, common.EVENT_TOPIC_NAME)
	collect.InitCollectors(s)
	common.Logger.Debug("server start")
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		common.Logger.Fatalf("error: %v", err)
	}
}
