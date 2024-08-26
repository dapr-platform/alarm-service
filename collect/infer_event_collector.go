package collect

import (
	"alarm-service/event_process"
	"alarm-service/model"
	"alarm-service/service"
	"context"
	"encoding/json"
	mycommon "github.com/dapr-platform/common"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"time"
)

var INFER_COLLECT_IDENTIFIER = "infer"

func NewInferEventHandler(server common.Service, eventPub string, eventTopic string) (err error) {
	var sub = &common.Subscription{
		PubsubName: eventPub,
		Topic:      eventTopic,
		Route:      "/inferEventHandler",
	}

	return server.AddTopicEventHandler(sub, inferEventHandler)
}
func parseEvent(data *model.Event) (eventAlarm *model.Alarm, err error) {
	eventAlarm = &model.Alarm{
		ID:          data.ID,
		Dn:          data.Dn,
		Title:       data.EventTitle,
		Type:        mycommon.EventTypePlatform,
		Description: data.EventText,
		Status:      data.Status,
		Level:       data.Level,
		EventTime:   data.EventTime,
		CreateAt:    mycommon.LocalTime(time.Now()),
		ObjectID:    data.ObjectID,
		ObjectName:  data.ObjectID, //TODO
		Location:    "infer-service",
		TotalCount:  1,
		ExtraData:   data.EventExtra,
	}
	return
}
func inferEventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s\n", e.PubsubName, e.Topic, e.ID, e.Data)
	var data = model.Event{}
	err = json.Unmarshal(e.RawData, &data)
	if err != nil {
		log.Println("event - unmarshal error: ", err) //TODO send exception alarm ?
		return false, nil
	}
	err = SaveEvent(ctx, data)
	if err != nil {
		log.Println("event save error ", err)
		return false, nil
	}
	err = SendEventToWeb(ctx, data)
	if err != nil {
		log.Println("event send error ", err)
	}
	alarmEvent, err := parseEvent(&data) //转成alarmService的event结构
	if err != nil {
		log.Println("event - parse error: ", err)
		return false, nil
	}
	if alarmEvent != nil {
		alarm, err := event_process.ProcessEventToAlarm(INFER_COLLECT_IDENTIFIER, alarmEvent) //根据event判断是否需要生成告警
		if err != nil {
			log.Println("event - process error: ", err)
			return false, nil
		}
		if alarm != nil {
			go service.ProcessEvent(context.Background(), alarm)
		}
	}

	return false, nil
}
