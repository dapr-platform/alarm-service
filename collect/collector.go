package collect

import (
	"alarm-service/entity"
	"alarm-service/model"
	"context"
	"encoding/json"
	mycommon "github.com/dapr-platform/common"
	"github.com/dapr/go-sdk/service/common"
)

func ListCollector() (collectors []entity.Collector) {
	collectors = make([]entity.Collector, 0)
	collectors[0] = entity.Collector{
		Identifier: INFER_COLLECT_IDENTIFIER,
		Name:       "推理事件采集",
	}
	return collectors
}
func InitCollectors(server common.Service) {
	err := NewInferEventHandler(server, mycommon.PUBSUB_NAME, "infer_events")
	if err != nil {
		panic(err)
	}
}
func SaveEvent(ctx context.Context, event model.Event) (err error) {
	return mycommon.DbUpsert[model.Event](ctx, mycommon.GetDaprClient(), event, model.EventTableInfo.Name, model.Event_FIELD_NAME_id)
}

func SendEventToWeb(ctx context.Context, event model.Event) (err error) {
	buf, err := json.Marshal(event)
	if err != nil {
		mycommon.Logger.Error(err.Error())
		return
	}
	msg := mycommon.CommonMessage{
		mycommon.COMMON_MESSAGE_KEY_MARK:       "event",
		mycommon.COMMON_MESSAGE_KEY_CONNECT_ID: "",
		mycommon.COMMON_MESSAGE_KEY_MESSAGE:    string(buf),
	}
	err = mycommon.GetDaprClient().PublishEvent(ctx, mycommon.PUBSUB_NAME, mycommon.WEB_MESSAGE_TOPIC_NAME, msg)
	if err != nil {
		mycommon.Logger.Error(err.Error())
		return
	}
	return
}
