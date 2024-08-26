package entity

type AlarmQuery struct {
	StartTime string `json:"start_time,omitempty"` // 查询开始时间 >,为空则为当天0点, 格式 yyyy-MM-dd HH:mm:ss
	EndTime   string `json:"end_time,omitempty"`   // 查询结束时间 <=，为空则为当前时间 格式 yyyy-MM-dd HH:mm:ss
	Types     string `json:"types,omitempty"`      // 告警类型，为""则为全部类型，以,分隔,例如：1,2,3
	SubTypes  string `json:"sub_types,omitempty"`  // 告警子类型，为"" 则为全部子类型，以,分隔,例如：1,2,3
	Levels    string `json:"levels,omitempty"`     // 告警级别，为"" 则为全部级别， 以,分隔,例如：1,2,3
	Status    string `json:"status,omitempty"`     // 告警状态，为"" 则为全部状态，以,分隔,例如：0,1, 0:已清除 1:未清除
	ObjectIds string `json:"object_ids,omitempty"` // 告警对象ID，为"" 则为全部对象，以,分隔,例如：1,2,3
}

type AlarmIdsReq struct {
	Ids string `json:"ids"` // 告警ID，以,分隔,例如：1,2,3

}
