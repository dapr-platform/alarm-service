package api

import (
	"alarm-service/entity"
	"alarm-service/model"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func InitCustomEventRouter(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/send-event", sendEventHandler)
	r.Post(common.BASE_CONTEXT+"/grafana-event", receiveGrafanaEventHandler)
	r.Post(common.BASE_CONTEXT+"/nightingale-event", receiveNightingaleEventHandler)
	r.Post(common.BASE_CONTEXT+"/gocron-event", receiveGoCronEventHandler)
}

// @Summary 发送事件
// @Description 发送事件,测试使用。时间格式为 2006-01-02T15:04:05
// @Tags Message
// @Param req body model.Alarm true "发送消息"
// @Produce  json
// @Success 200 {object} common.Response "成功消息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /send-event [post]
func sendEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		common.HttpError(w, common.ErrReqBodyRead.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	req := &model.Alarm{}
	err = json.Unmarshal(body, req)
	if err != nil {
		common.HttpError(w, common.ErrReqBodyParse.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Println("send event:", req)
	err = common.GetDaprClient().PublishEvent(r.Context(), common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, req)
	if err != nil {
		common.HttpError(w, common.ErrPubSubPublish.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}

	common.HttpSuccess(w, nil)
}

// @Summary 接收nightingale事件
// @Description 接收nightingale事件
// @Tags Message
// @Param req body string true "消息"
// @Produce  json
// @Success 200 {object} common.Response "成功消息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /nightingale-event [post]
func receiveNightingaleEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		common.HttpError(w, common.ErrReqBodyRead.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Println(string(body))

	common.HttpSuccess(w, common.OK)
}

// @Summary 接收grafana事件
// @Description 接收grafana事件
// @Tags Message
// @Param req body string true "消息"
// @Produce  json
// @Success 200 {object} common.Response "成功消息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /grafana-event [post]
func receiveGrafanaEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		common.HttpError(w, common.ErrReqBodyRead.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Println(string(body))
	var notification entity.GrafanaNotification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		common.HttpError(w, common.ErrReqBodyParse.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	for _, alert := range notification.Alerts {
		isAlarm := common.EventStatusClosed
		if alert.Status == "firing" {
			isAlarm = common.EventStatusActive
		}
		level := common.EventLevelMinor
		if val, exists := alert.Annotations["Level"]; exists {
			v, _ := strconv.Atoi(val)
			level = int32(v)
		}
		var event = model.Alarm{
			Dn:          "GRAFANA:" + alert.Fingerprint,
			Title:       alert.Labels.Alertname,
			Status:      isAlarm,
			Type:        common.EventTypePlatform,
			Level:       level,
			EventTime:   common.LocalTime(time.Now()),
			Description: alert.ValueString,
		}
		err = common.GetDaprClient().PublishEvent(r.Context(), common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, event)
		if err != nil {
			common.HttpError(w, common.ErrPubSubPublish.AppendMsg(err.Error()), http.StatusInternalServerError)
			return
		}
	}

	common.HttpSuccess(w, nil)
}

// @Summary 接收gocron事件
// @Description 接收gocron事件
// @Tags Message
// @Param req body entity.GoCronEvent true "消息"
// @Produce  json
// @Success 200 {object} common.Response "成功消息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /gocron-event [post]
func receiveGoCronEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		common.HttpError(w, common.ErrReqBodyRead.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Println(string(body))
	var gocronEvent entity.GoCronEvent
	err = json.Unmarshal(body, &gocronEvent)
	if err != nil {
		common.HttpError(w, common.ErrReqBodyParse.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	isAlarm := common.EventStatusClosed
	if gocronEvent.Status == "失败" {
		isAlarm = common.EventStatusActive
	}

	var event = model.Alarm{
		Dn:          "GOCRON:" + gocronEvent.Name,
		Title:       gocronEvent.Name + gocronEvent.Status,
		Status:      isAlarm,
		Type:        common.EventTypePlatform,
		Level:       common.EventLevelMinor,
		EventTime:   common.LocalTime(time.Now()),
		Description: gocronEvent.Result,
	}
	err = common.GetDaprClient().PublishEvent(r.Context(), common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, event)
	if err != nil {
		common.HttpError(w, common.ErrPubSubPublish.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}

	common.HttpSuccess(w, nil)
}
