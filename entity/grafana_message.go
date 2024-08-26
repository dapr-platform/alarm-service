package entity

import "time"

type Alert struct {
	Status       string            `json:"status"`
	Labels       AlertLabels       `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL"`
	PanelURL     string            `json:"panelURL"`
	ValueString  string            `json:"valueString"`
}
type AlertLabels struct {
	Alertname string `json:"alertname"`
}

type GrafanaNotification struct {
	Receiver          string                 `json:"receiver"`
	Status            string                 `json:"status"`
	Alerts            []Alert                `json:"alerts"`
	GroupLabels       map[string]interface{} `json:"groupLabels"`
	CommonLabels      map[string]interface{} `json:"commonLabels"`
	CommonAnnotations map[string]interface{} `json:"commonAnnotations"`
	ExternalURL       string                 `json:"externalURL"`
	Version           string                 `json:"version"`
	GroupKey          string                 `json:"groupKey"`
	TruncatedAlerts   int                    `json:"truncatedAlerts"`
	OrgId             int                    `json:"orgId"`
	Title             string                 `json:"title"`
	State             string                 `json:"state"`
	Message           string                 `json:"message"`
}
