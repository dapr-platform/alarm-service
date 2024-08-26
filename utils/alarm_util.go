package utils

import (
	"alarm-service/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dapr-platform/common"
)

func GetActiveAlarmFromDB(ctx context.Context, dn string) (alarm *model.Alarm, err error) {
	ret, err := common.DbQuery[model.Alarm](ctx, common.GetDaprClient(), "f_alarm", "dn="+dn+"&status=1")
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, nil
	}
	a := ret[0]
	return &a, nil

}

func JsonInt64ToString(vt interface{}) string {
	st := ""
	switch vt.(type) {
	case json.Number:
		st = vt.(json.Number).String()
	case float64:
		st = fmt.Sprintf("%.0f", vt.(float64))
	case float32:
		st = fmt.Sprintf("%.0f", vt.(float32))
	case int64:
		st = fmt.Sprintf("%d", vt.(int64))
	case int32:
		st = fmt.Sprintf("%d", vt.(int32))
	case int:
		st = fmt.Sprintf("%d", vt.(int))
	case string:
		st = vt.(string)
	}
	return st
}
