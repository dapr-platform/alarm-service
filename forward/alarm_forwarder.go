package forward

import (
	"alarm-service/entity"
	"alarm-service/model"
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"log"
	"strings"
)

type AlarmForwarder interface {
	Name() string
	Forward(ctx context.Context, alarm *model.Alarm) error
}

var AND_TYPE = 0
var OR_TYPE = 1

type TopicAlarmForwarder struct {
	Topic     string
	AndOrType int // 0: and, 1: or
	Filters   []entity.Filter
}

var AlarmForwarders []AlarmForwarder

func init() {
	f := &TopicAlarmForwarder{
		Topic:     "web",
		AndOrType: OR_TYPE,
		Filters:   []entity.Filter{},
	}
	/*
		f := &TopicAlarmForwarder{
			Topic:     "alarm",
			AndOrType: OR_TYPE,
			Filters: []entity.Filter{entity.Filter{
				Items: []entity.FilterFunc{&entity.FilterItem[float64]{
					Name:  "type",
					Op:    "eq",
					Value: 1,
				},
				},
			},
			},
		}

	*/

	AlarmForwarders = append(AlarmForwarders, f)
	//tf := &TicketAlarmForwarder{}
	//AlarmForwarders = append(AlarmForwarders, tf)
}
func (t *TopicAlarmForwarder) Name() string {
	return t.Topic
}
func (t *TopicAlarmForwarder) Forward(ctx context.Context, alarm *model.Alarm) error {
	buf, _ := json.Marshal(alarm)

	var mapAlarm map[string]interface{}
	json.Unmarshal(buf, &mapAlarm)
	//javascript int64范围不够，需要转换成字符串

	//mapAlarm["id"] = utils.JsonInt64ToString(mapAlarm["id"])
	//TransTimestampFieldToFormatString(mapAlarm)
	if len(t.Filters) > 0 {
		allFilters := make([]bool, 0)
		for _, filter := range t.Filters {
			oneFilter := true
			for _, item := range filter.Items {
				if !item.Filter(mapAlarm) {
					oneFilter = false
					break
				}
			}
			allFilters = append(allFilters, oneFilter)
		}
		if t.AndOrType == OR_TYPE {
			flag := false
			for _, oneFilter := range allFilters {
				if oneFilter {
					flag = true
					break
				}
			}
			if !flag {
				return nil
			}
		} else {
			for _, oneFilter := range allFilters {
				if !oneFilter {
					return nil
				}
			}
		}
	}
	alarmStr, _ := json.Marshal(mapAlarm)
	msg := common.CommonMessage{
		common.COMMON_MESSAGE_KEY_MARK:    "alarm",
		common.COMMON_MESSAGE_KEY_MESSAGE: string(alarmStr),
	}
	err := common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, t.Topic, msg)
	log.Println("publish alarm:", t.Topic, alarm.ID, "err=", err)
	return err
}

func TransTimestampFieldToFormatString(mapAlarm map[string]interface{}) {
	fields := []string{"event_time", "clear_time", "ack_time", "create_at", "update_at"}
	for _, field := range fields {
		if value, ok := mapAlarm[field]; ok {
			if value != nil {
				mapAlarm[field] = strings.Replace(value.(string), "T", " ", -1)
			}
		}
	}

}
