package api

import (
	"alarm-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitCollector_scriptRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/collector-script/page", Collector_scriptPageListHandler)
	r.Get(common.BASE_CONTEXT+"/collector-script", Collector_scriptListHandler)
	r.Post(common.BASE_CONTEXT+"/collector-script", UpsertCollector_scriptHandler)
	r.Delete(common.BASE_CONTEXT+"/collector-script/{id}", DeleteCollector_scriptHandler)
	r.Post(common.BASE_CONTEXT+"/collector-script/batch-delete", batchDeleteCollector_scriptHandler)
	r.Post(common.BASE_CONTEXT+"/collector-script/batch-upsert", batchUpsertCollector_scriptHandler)
	r.Get(common.BASE_CONTEXT+"/collector-script/groupby", Collector_scriptGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Collector_script
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /collector-script/groupby [get]
func Collector_scriptGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_collector_script")
}

// @Summary batch update
// @Description batch update
// @Tags Collector_script
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /collector-script/batch-upsert [post]
func batchUpsertCollector_scriptHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Collector_scriptTableInfo.Name, model.Collector_script_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Collector_script
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param script query string false "script"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Collector_script}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /collector-script/page [get]
func Collector_scriptPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Collector_script](w, r, common.GetDaprClient(), "o_collector_script", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Collector_script
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param script query string false "script"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Collector_script} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /collector-script [get]
func Collector_scriptListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Collector_script](w, r, common.GetDaprClient(), "o_collector_script", "id")
}

// @Summary save
// @Description save
// @Tags Collector_script
// @Accept       json
// @Param item body model.Collector_script true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Collector_script} "object"
// @Failure 500 {object} common.Response ""
// @Router /collector-script [post]
func UpsertCollector_scriptHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Collector_script
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Collector_script")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Collector_script)
	}

	err = common.DbUpsert[model.Collector_script](r.Context(), common.GetDaprClient(), val, model.Collector_scriptTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Collector_script
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Collector_script} "object"
// @Failure 500 {object} common.Response ""
// @Router /collector-script/{id} [delete]
func DeleteCollector_scriptHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Collector_script")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_collector_script", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Collector_script
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /collector-script/batch-delete [post]
func batchDeleteCollector_scriptHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Collector_script")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_collector_script", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
