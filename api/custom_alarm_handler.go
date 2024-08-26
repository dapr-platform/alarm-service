package api

import (
	"alarm-service/entity"
	"alarm-service/service"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"log"
	"net/http"
	"strconv"
)

var (
	_ = null.String{}
)

func InitCustomAlarmRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/alarm-query", alarmQueryHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-clear", alarmClearHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-ack", alarmAckHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-archive", alarmArchiveHandler)

}

// @Summary 查询告警
// @Description 查询告警
// @Tags Alarm
// @Param _page query int false "current page, default is 1"
// @Param _page_size query int false "page size,default is 10"
// @param alarm_query body entity.AlarmQuery true "query condition"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Alarm}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /alarm-query [post]
func alarmQueryHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}
	var alarmQuery entity.AlarmQuery
	if err := common.ReadRequestBody(r, &alarmQuery); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("read request body error: "+err.Error()))
		return
	}
	pageI, err := strconv.Atoi(page)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("page must be int"))
		return
	}
	pageSizeI, err := strconv.Atoi(pageSize)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("page_size must be int"))
		return
	}
	s, _ := json.Marshal(alarmQuery)
	log.Println("alarmQuery:", string(s))
	ret, err := service.QueryAlarm(r.Context(), pageI, pageSizeI, &alarmQuery)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	buf, _ := json.Marshal(ret)
	log.Println("alarmQuery ret:", string(buf))
	common.HttpSuccess(w, common.OK.WithData(ret))

}

// @Summary 告警清除
// @Description 告警清除
// @Tags Alarm
// @param alarm_clear body entity.AlarmIdsReq true "ids"
// @Produce  json
// @Success 200 {object} common.Response "{status:0}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /alarm-clear [post]
func alarmClearHandler(w http.ResponseWriter, r *http.Request) {

	var alarmClearReq entity.AlarmIdsReq
	if err := common.ReadRequestBody(r, &alarmClearReq); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("alarm clear read request body error: "+err.Error()))
		return
	}
	buf, _ := json.Marshal(alarmClearReq)
	log.Println("alarmClearReq:", string(buf))
	sub, _ := common.ExtractUserSub(r)
	err := service.ClearAlarm(r.Context(), sub, &alarmClearReq)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)

}

// @Summary 告警归档
// @Description 告警归档
// @Tags Alarm
// @param alarm_clear body entity.AlarmIdsReq true "ids"
// @Produce  json
// @Success 200 {object} common.Response "{status:0}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /alarm-archive [post]
func alarmArchiveHandler(w http.ResponseWriter, r *http.Request) {

	var alarmIdsReq entity.AlarmIdsReq
	if err := common.ReadRequestBody(r, &alarmIdsReq); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("alarm archive read request body error: "+err.Error()))
		return
	}
	buf, _ := json.Marshal(alarmIdsReq)
	log.Println("alarmArchiveReq:", string(buf))
	sub, _ := common.ExtractUserSub(r)
	err := service.ArchiveAlarm(r.Context(), sub, &alarmIdsReq)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)

}

// @Summary 告警确认
// @Description 告警确认
// @Tags Alarm
// @param alarm_ids body entity.AlarmIdsReq true "ids"
// @Produce  json
// @Success 200 {object} common.Response "{status:0}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /alarm-ack [post]
func alarmAckHandler(w http.ResponseWriter, r *http.Request) {

	var alarmIdsReq entity.AlarmIdsReq
	if err := common.ReadRequestBody(r, &alarmIdsReq); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("alarm clear read request body error: "+err.Error()))
		return
	}
	sub, _ := common.ExtractUserSub(r)
	buf, _ := json.Marshal(alarmIdsReq)
	log.Println("alarmAckReq:", string(buf))
	err := service.AckAlarm(r.Context(), sub, &alarmIdsReq)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)

}
