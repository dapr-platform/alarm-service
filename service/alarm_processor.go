package service

import (
	"alarm-service/forward"
	"alarm-service/model"
	"alarm-service/prom"
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"log"
	"strings"
	"time"
)

func ProcessEvent(ctx context.Context, event *model.Alarm) {
	alarm, err := preProcess(ctx, event)
	if err != nil {
		log.Println("Error in preProcessing alarm: ", err)
		return
	}
	if alarm == nil {
		return
	}
	go prom.CountAlarm(cast.ToString(alarm.Level))
	forwardTo(ctx, alarm)

	err = saveToDb(ctx, alarm)
	if err != nil {
		log.Println("Error in saving alarm to db: ", err)
		log.Println("Alarm: ", alarm)
		return
	}
	log.Println("Alarm saved to db,id: ", alarm.ID)
}

func preProcess(ctx context.Context, event *model.Alarm) (alarm *model.Alarm, err error) {
	log.Println("Processing event: ", event)
	alarm, err = common.DbGetOne[model.Alarm](ctx, common.GetDaprClient(), model.AlarmTableInfo.Name, "dn="+event.Dn)
	if err != nil {
		log.Println("Error in getting alarm from db: ", err)
		err = errors.Wrap(err, "Error in getting alarm from db")
		return

	}
	if alarm != nil {
		t := common.LocalTime(time.Now())
		alarm.UpdatedAt = t
		alarm.EventTime = t
		alarm.Title = event.Title
		alarm.Description = event.Description
		alarm.Level = event.Level
		alarm.Status = event.Status
		alarm.Location = event.Location
		alarm.Type = event.Type
		alarm.AckUser = event.AckUser
		alarm.AckFlag = event.AckFlag
		alarm.AckTime = event.AckTime
		alarm.ArchiveTime = event.ArchiveTime
		alarm.ArchiveFlag = event.ArchiveFlag
		alarm.ArchiveUser = event.ArchiveUser
		if event.Status == common.EventStatusClosed {
			if alarm.ClearUser == "" {
				alarm.ClearUser = "system"
			}
			alarm.ClearTime = common.LocalTime(time.Now())
			return
		} else {
			alarm.TotalCount = alarm.TotalCount + 1
			return
		}
	} else {
		if event.Status == common.EventStatusClosed {
			return
		}
		alarm = &model.Alarm{}
		alarm.ID = common.NanoId()
		alarm.Dn = event.Dn
		alarm.Title = event.Title
		alarm.Description = event.Description
		alarm.Level = event.Level
		alarm.Type = event.Type
		alarm.Status = event.Status
		alarm.EventTime = event.EventTime
		alarm.ObjectID = event.ObjectID
		alarm.ObjectName = event.ObjectName
		alarm.Location = event.Location
		alarm.CreateAt = common.LocalTime(time.Now())
		alarm.UpdatedAt = common.LocalTime(time.Now())
		alarm.TotalCount = 1
		return
	}
	return
}

func forwardTo(ctx context.Context, alarm *model.Alarm) {
	for _, forwarder := range forward.AlarmForwarders {
		err := forwarder.Forward(ctx, alarm)
		if err != nil {
			log.Println("Error in forwarding alarm to ", forwarder.Name, ":", err)
			continue
		}
	}

}

func saveToDb(ctx context.Context, alarm *model.Alarm) (err error) {
	escapeDbField(alarm)

	if alarm.ArchiveFlag == common.EventArchivedFlag {

		_, err = common.DbInsert[model.Alarm](ctx, common.GetDaprClient(), *alarm, "f_history_alarm")
		if err != nil {
			log.Println("insert history alarm error " + alarm.Dn)
			return err
		}

		err = common.DbDelete(ctx, common.GetDaprClient(), "f_alarm", "id", alarm.ID)
		if err != nil {
			log.Println("delete alarm error " + alarm.Dn)
			return err
		}
	} else {
		return common.DbUpsert[model.Alarm](ctx, common.GetDaprClient(), *alarm, "f_alarm", "dn")
	}
	return nil
}
func escapeDbField(alarm *model.Alarm) {
	alarm.Title = strings.Replace(alarm.Title, "'", "''", -1)
	alarm.Description = strings.Replace(alarm.Description, "'", "''", -1)
}
