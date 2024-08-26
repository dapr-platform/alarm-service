package event_process

import "alarm-service/model"

func ProcessEventToAlarm(collectorIdentifier string, event *model.Alarm) (alarm *model.Alarm, err error) {
	return event, nil
}
