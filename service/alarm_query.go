package service

import (
	"alarm-service/entity"
	"alarm-service/model"
	"alarm-service/utils"
	"context"
	"github.com/dapr-platform/common"
	"log"
	"strings"
	"time"
)

func ClearAlarm(ctx context.Context, sub string, req *entity.AlarmIdsReq) error {

	for _, id := range strings.Split(req.Ids, ",") {
		existAlarm, err := common.DbGetOne[model.Alarm](ctx, common.GetDaprClient(), "f_alarm", "id="+id)
		if err != nil {
			log.Println("query alarm error:", id, err)
			return err
		}
		if existAlarm == nil {
			log.Println("alarm not exist:", id)
			continue
		}

		existAlarm.Status = common.EventStatusClosed
		existAlarm.ClearUser = sub
		err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, existAlarm)
		if err != nil {
			log.Println("publish event error:", err)
			return err
		}
		log.Println("send clear alarm event:", id)
	}
	return nil
}

// 发到event topic。
func AckAlarm(ctx context.Context, user string, req *entity.AlarmIdsReq) error {

	for _, id := range strings.Split(req.Ids, ",") {
		existAlarm, err := common.DbGetOne[model.Alarm](ctx, common.GetDaprClient(), "f_alarm", "id="+id)
		if err != nil {
			log.Println("query alarm error:", id, err)
			return err
		}
		if existAlarm == nil {
			log.Println("alarm not exist:", id)
			continue
		}

		existAlarm.AckFlag = 1
		existAlarm.AckTime = common.LocalTime(time.Now())
		existAlarm.AckUser = user

		err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, existAlarm)
		if err != nil {
			log.Println("publish event error:", err)
			return err
		}
		log.Println("send ack alarm event:", id)
	}
	return nil
}
func ArchiveAlarm(ctx context.Context, user string, req *entity.AlarmIdsReq) error {

	for _, id := range strings.Split(req.Ids, ",") {
		existAlarm, err := common.DbGetOne[model.Alarm](ctx, common.GetDaprClient(), "f_alarm", "id="+id)
		if err != nil {
			log.Println("query alarm error:", id, err)
			return err
		}
		if existAlarm == nil {
			log.Println("alarm not exist:", id)
			continue
		}

		existAlarm.ArchiveFlag = 1
		existAlarm.ArchiveTime = common.LocalTime(time.Now())
		existAlarm.ArchiveUser = user

		err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, common.EVENT_TOPIC_NAME, existAlarm)
		if err != nil {
			log.Println("publish event error:", err)
			return err
		}
		log.Println("send archive alarm event:", id)
	}
	return nil
}
func QueryAlarm(ctx context.Context, page, pageSize int, query *entity.AlarmQuery) (*common.PageGeneric[map[string]any], error) {
	var queryPara = ""
	startTimeStr := query.StartTime
	if startTimeStr == "" {
		dayZero := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
		startTimeStr = dayZero.Format("2006-01-02T15:04:05")
	} else {
		startTimeStr = strings.Replace(query.StartTime, " ", "T", 1)
	}
	queryPara += "event_time='$gte." + startTimeStr + "'"
	endTimeStr := query.EndTime
	if endTimeStr == "" {
		endTimeStr = time.Now().Format("2006-01-02T15:04:05")
	} else {
		endTimeStr = strings.Replace(query.EndTime, " ", "T", 1)
	}
	queryPara += "&event_time='$lt." + endTimeStr + "'"

	if query.Types != "" {
		queryPara += "&type=$in." + query.Types
	}

	if query.SubTypes != "" {

		queryPara += "&subtype=$in." + query.SubTypes
	}

	if query.Levels != "" {

		queryPara += "&level=$in." + query.Levels
	}
	if query.Status != "" {

		queryPara += "&status=$in." + query.Status
	}
	if query.ObjectIds != "" {

		queryPara += "&object_id=$in." + query.ObjectIds + ""
	}
	data, err := common.DbPageQuery[map[string]any](ctx, common.GetDaprClient(), page, pageSize, "-updated_at", "f_alarm", "id", queryPara)
	if err != nil {
		return nil, err
	}
	//javascript int64范围不够，需要转换成字符串
	for _, v := range data.Items {
		v["id"] = utils.JsonInt64ToString(v["id"])
	}
	return data, err
}
