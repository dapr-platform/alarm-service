package api

import (
	"alarm-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitHistory_alarmRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/history-alarm/page", History_alarmPageListHandler)
	r.Get(common.BASE_CONTEXT+"/history-alarm", History_alarmListHandler)
	r.Post(common.BASE_CONTEXT+"/history-alarm", UpsertHistory_alarmHandler)
	r.Delete(common.BASE_CONTEXT+"/history-alarm/{id}", DeleteHistory_alarmHandler)
	r.Post(common.BASE_CONTEXT+"/history-alarm/batch-delete", batchDeleteHistory_alarmHandler)
	r.Post(common.BASE_CONTEXT+"/history-alarm/batch-upsert", batchUpsertHistory_alarmHandler)
	r.Get(common.BASE_CONTEXT+"/history-alarm/groupby", History_alarmGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags History_alarm
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /history-alarm/groupby [get]
func History_alarmGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "f_history_alarm")
}

// @Summary batch update
// @Description batch update
// @Tags History_alarm
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /history-alarm/batch-upsert [post]
func batchUpsertHistory_alarmHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.History_alarmTableInfo.Name, model.History_alarm_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags History_alarm
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param type query string false "type"
// @Param description query string false "description"
// @Param status query string false "status"
// @Param level query string false "level"
// @Param event_time query string false "event_time"
// @Param create_at query string false "create_at"
// @Param updated_at query string false "updated_at"
// @Param object_id query string false "object_id"
// @Param object_name query string false "object_name"
// @Param location query string false "location"
// @Param total_count query string false "total_count"
// @Param clear_time query string false "clear_time"
// @Param ack_flag query string false "ack_flag"
// @Param archive_flag query string false "archive_flag"
// @Param ack_time query string false "ack_time"
// @Param clear_user query string false "clear_user"
// @Param ack_user query string false "ack_user"
// @Param archive_time query string false "archive_time"
// @Param archive_user query string false "archive_user"
// @Param extra_data query string false "extra_data"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.History_alarm}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /history-alarm/page [get]
func History_alarmPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.History_alarm](w, r, common.GetDaprClient(), "f_history_alarm", "id")

}

// @Summary query objects
// @Description query objects
// @Tags History_alarm
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param type query string false "type"
// @Param description query string false "description"
// @Param status query string false "status"
// @Param level query string false "level"
// @Param event_time query string false "event_time"
// @Param create_at query string false "create_at"
// @Param updated_at query string false "updated_at"
// @Param object_id query string false "object_id"
// @Param object_name query string false "object_name"
// @Param location query string false "location"
// @Param total_count query string false "total_count"
// @Param clear_time query string false "clear_time"
// @Param ack_flag query string false "ack_flag"
// @Param archive_flag query string false "archive_flag"
// @Param ack_time query string false "ack_time"
// @Param clear_user query string false "clear_user"
// @Param ack_user query string false "ack_user"
// @Param archive_time query string false "archive_time"
// @Param archive_user query string false "archive_user"
// @Param extra_data query string false "extra_data"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.History_alarm} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /history-alarm [get]
func History_alarmListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.History_alarm](w, r, common.GetDaprClient(), "f_history_alarm", "id")
}

// @Summary save
// @Description save
// @Tags History_alarm
// @Accept       json
// @Param item body model.History_alarm true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.History_alarm} "object"
// @Failure 500 {object} common.Response ""
// @Router /history-alarm [post]
func UpsertHistory_alarmHandler(w http.ResponseWriter, r *http.Request) {
	var val model.History_alarm
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("History_alarm")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.History_alarm)
	}

	err = common.DbUpsert[model.History_alarm](r.Context(), common.GetDaprClient(), val, model.History_alarmTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags History_alarm
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.History_alarm} "object"
// @Failure 500 {object} common.Response ""
// @Router /history-alarm/{id} [delete]
func DeleteHistory_alarmHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("History_alarm")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "f_history_alarm", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags History_alarm
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /history-alarm/batch-delete [post]
func batchDeleteHistory_alarmHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("History_alarm")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "f_history_alarm", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
