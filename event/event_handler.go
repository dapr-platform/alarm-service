package event

import (
	"alarm-service/model"
	"alarm-service/service"
	"context"
	"encoding/json"
	"github.com/dapr/go-sdk/service/common"
	"log"
)

func NewEventHandler(server common.Service, eventPub string, eventTopic string) {
	var sub = &common.Subscription{
		PubsubName: eventPub,
		Topic:      eventTopic,
		Route:      "/eventHandler",
	}

	err := server.AddTopicEventHandler(sub, eventHandler)

	if err != nil {
		panic(err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s\n", e.PubsubName, e.Topic, e.ID, e.Data)
	var event = &model.Alarm{}
	err = json.Unmarshal(e.RawData, event)
	if err != nil {
		log.Println("event - unmarshal error: ", err)
	} else {
		go service.ProcessEvent(context.Background(), event)
	}

	return false, nil
}
