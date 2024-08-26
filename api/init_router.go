package api

import (
	"github.com/go-chi/chi/v5"
)

func InitRoute(r chi.Router) {
	InitAlarmRoute(r)
	InitHistory_alarmRoute(r)
	InitCustomEventRouter(r)
	InitCustomAlarmRoute(r)
	InitCustomEventRouter(r)
	InitCollector_scriptRoute(r)
	InitEventRoute(r)
}
