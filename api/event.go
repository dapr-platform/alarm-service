package api

import (
	"alarm-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitEventRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/event/page", EventPageListHandler)
	r.Get(common.BASE_CONTEXT+"/event", EventListHandler)
	r.Post(common.BASE_CONTEXT+"/event", UpsertEventHandler)
	r.Delete(common.BASE_CONTEXT+"/event/{id}", DeleteEventHandler)
	r.Post(common.BASE_CONTEXT+"/event/batch-delete", batchDeleteEventHandler)
	r.Post(common.BASE_CONTEXT+"/event/batch-upsert", batchUpsertEventHandler)
	r.Get(common.BASE_CONTEXT+"/event/groupby", EventGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Event
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /event/groupby [get]
func EventGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "f_event")
}

// @Summary batch update
// @Description batch update
// @Tags Event
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /event/batch-upsert [post]
func batchUpsertEventHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.EventTableInfo.Name, model.Event_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Event
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param event_title query string false "event_title"
// @Param event_text query string false "event_text"
// @Param event_time query string false "event_time"
// @Param object_id query string false "object_id"
// @Param status query string false "status"
// @Param event_extra query string false "event_extra"
// @Param level query string false "level"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Event}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /event/page [get]
func EventPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Event](w, r, common.GetDaprClient(), "f_event", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Event
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param event_title query string false "event_title"
// @Param event_text query string false "event_text"
// @Param event_time query string false "event_time"
// @Param object_id query string false "object_id"
// @Param status query string false "status"
// @Param event_extra query string false "event_extra"
// @Param level query string false "level"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Event} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /event [get]
func EventListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Event](w, r, common.GetDaprClient(), "f_event", "id")
}

// @Summary save
// @Description save
// @Tags Event
// @Accept       json
// @Param item body model.Event true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Event} "object"
// @Failure 500 {object} common.Response ""
// @Router /event [post]
func UpsertEventHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Event
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Event")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Event)
	}

	err = common.DbUpsert[model.Event](r.Context(), common.GetDaprClient(), val, model.EventTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Event
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Event} "object"
// @Failure 500 {object} common.Response ""
// @Router /event/{id} [delete]
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Event")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "f_event", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Event
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /event/batch-delete [post]
func batchDeleteEventHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Event")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "f_event", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
