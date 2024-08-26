package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: f_history_alarm
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] dn                                             VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] title                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] description                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] status                                         INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 6] level                                          INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 7] event_time                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 8] create_at                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 9] updated_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[10] object_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] object_name                                    VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] location                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] total_count                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[14] clear_time                                     TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[15] ack_flag                                       INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[16] archive_flag                                   INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[17] ack_time                                       TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[18] clear_user                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[19] ack_user                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[20] archive_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[21] archive_user                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[22] extra_data                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "QfEtdRsamikJwUfQnhWITNRfE",    "dn": "uJCgDojIgCXaJQQiXHcRfgFOl",    "title": "mWJaYGbcKHBeLnJsYVQCCBeng",    "type": 15,    "description": "QBoqgnOYWBvHIIECJlxBSglEX",    "status": 8,    "level": 27,    "event_time": 91,    "create_at": 83,    "updated_at": 12,    "object_id": "JJPVDcnthXZFyZCReCsRLuNZi",    "object_name": "dEHFvWikONXDFfVQlrrEilYux",    "location": "dDFadgnxDOFXDRjRURwEsoYYV",    "total_count": 70,    "clear_time": 96,    "ack_flag": 71,    "archive_flag": 23,    "ack_time": 63,    "clear_user": "noiIVEXfOSCToSIIXixOfdBJQ",    "ack_user": "ONPbjEWKCfZCPWifewLHfNoun",    "archive_time": 88,    "archive_user": "YymlHPfuGHbDUrOYUwWHBONTZ",    "extra_data": "jOwvKltcHehoSdOroiyUAZCxi"}



*/

var (
	History_alarm_FIELD_NAME_id = "id"

	History_alarm_FIELD_NAME_dn = "dn"

	History_alarm_FIELD_NAME_title = "title"

	History_alarm_FIELD_NAME_type = "type"

	History_alarm_FIELD_NAME_description = "description"

	History_alarm_FIELD_NAME_status = "status"

	History_alarm_FIELD_NAME_level = "level"

	History_alarm_FIELD_NAME_event_time = "event_time"

	History_alarm_FIELD_NAME_create_at = "create_at"

	History_alarm_FIELD_NAME_updated_at = "updated_at"

	History_alarm_FIELD_NAME_object_id = "object_id"

	History_alarm_FIELD_NAME_object_name = "object_name"

	History_alarm_FIELD_NAME_location = "location"

	History_alarm_FIELD_NAME_total_count = "total_count"

	History_alarm_FIELD_NAME_clear_time = "clear_time"

	History_alarm_FIELD_NAME_ack_flag = "ack_flag"

	History_alarm_FIELD_NAME_archive_flag = "archive_flag"

	History_alarm_FIELD_NAME_ack_time = "ack_time"

	History_alarm_FIELD_NAME_clear_user = "clear_user"

	History_alarm_FIELD_NAME_ack_user = "ack_user"

	History_alarm_FIELD_NAME_archive_time = "archive_time"

	History_alarm_FIELD_NAME_archive_user = "archive_user"

	History_alarm_FIELD_NAME_extra_data = "extra_data"
)

// History_alarm struct is a row record of the f_history_alarm table in the  database
type History_alarm struct {
	ID          string           `json:"id"`           //id
	Dn          string           `json:"dn"`           //dn
	Title       string           `json:"title"`        //告警标题
	Type        int32            `json:"type"`         //告警类型
	Description string           `json:"description"`  //告警描述
	Status      int32            `json:"status"`       //0:已清除告警,1:未清除告警
	Level       int32            `json:"level"`        //级别
	EventTime   common.LocalTime `json:"event_time"`   //告警时间
	CreateAt    common.LocalTime `json:"create_at"`    //创建时间
	UpdatedAt   common.LocalTime `json:"updated_at"`   //更新时间
	ObjectID    string           `json:"object_id"`    //告警对象id
	ObjectName  string           `json:"object_name"`  //告警对象名称
	Location    string           `json:"location"`     //告警位置
	TotalCount  int32            `json:"total_count"`  //发生次数
	ClearTime   common.LocalTime `json:"clear_time"`   //清除时间
	AckFlag     int32            `json:"ack_flag"`     //是否确认
	ArchiveFlag int32            `json:"archive_flag"` //是否归档，归档后才转移到历史告警
	AckTime     common.LocalTime `json:"ack_time"`     //确认时间
	ClearUser   string           `json:"clear_user"`   //清除人
	AckUser     string           `json:"ack_user"`     //确认人
	ArchiveTime common.LocalTime `json:"archive_time"` //归档时间
	ArchiveUser string           `json:"archive_user"` //归档人id
	ExtraData   string           `json:"extra_data"`   //extra_data

}

var History_alarmTableInfo = &TableInfo{
	Name: "f_history_alarm",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dn",
			Comment:            `dn`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Dn",
			GoFieldType:        "string",
			JSONFieldName:      "dn",
			ProtobufFieldName:  "dn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
			Comment:            `告警标题`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            `告警类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "description",
			Comment:            `告警描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `0:已清除告警,1:未清除告警`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "level",
			Comment:            `级别`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "event_time",
			Comment:            `告警时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "EventTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "event_time",
			ProtobufFieldName:  "event_time",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "create_at",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "create_at",
			ProtobufFieldName:  "create_at",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "updated_at",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "object_id",
			Comment:            `告警对象id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ObjectID",
			GoFieldType:        "string",
			JSONFieldName:      "object_id",
			ProtobufFieldName:  "object_id",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "object_name",
			Comment:            `告警对象名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ObjectName",
			GoFieldType:        "string",
			JSONFieldName:      "object_name",
			ProtobufFieldName:  "object_name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "location",
			Comment:            `告警位置`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Location",
			GoFieldType:        "string",
			JSONFieldName:      "location",
			ProtobufFieldName:  "location",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "total_count",
			Comment:            `发生次数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "TotalCount",
			GoFieldType:        "int32",
			JSONFieldName:      "total_count",
			ProtobufFieldName:  "total_count",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "clear_time",
			Comment:            `清除时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "ClearTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "clear_time",
			ProtobufFieldName:  "clear_time",
			ProtobufType:       "uint64",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ack_flag",
			Comment:            `是否确认`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "AckFlag",
			GoFieldType:        "int32",
			JSONFieldName:      "ack_flag",
			ProtobufFieldName:  "ack_flag",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "archive_flag",
			Comment:            `是否归档，归档后才转移到历史告警`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "ArchiveFlag",
			GoFieldType:        "int32",
			JSONFieldName:      "archive_flag",
			ProtobufFieldName:  "archive_flag",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "ack_time",
			Comment:            `确认时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "AckTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "ack_time",
			ProtobufFieldName:  "ack_time",
			ProtobufType:       "uint64",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "clear_user",
			Comment:            `清除人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ClearUser",
			GoFieldType:        "string",
			JSONFieldName:      "clear_user",
			ProtobufFieldName:  "clear_user",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "ack_user",
			Comment:            `确认人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AckUser",
			GoFieldType:        "string",
			JSONFieldName:      "ack_user",
			ProtobufFieldName:  "ack_user",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "archive_time",
			Comment:            `归档时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "ArchiveTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "archive_time",
			ProtobufFieldName:  "archive_time",
			ProtobufType:       "uint64",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "archive_user",
			Comment:            `归档人id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ArchiveUser",
			GoFieldType:        "string",
			JSONFieldName:      "archive_user",
			ProtobufFieldName:  "archive_user",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "extra_data",
			Comment:            `extra_data`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ExtraData",
			GoFieldType:        "string",
			JSONFieldName:      "extra_data",
			ProtobufFieldName:  "extra_data",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},
	},
}

// TableName sets the insert table name for this struct type
func (h *History_alarm) TableName() string {
	return "f_history_alarm"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *History_alarm) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *History_alarm) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *History_alarm) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *History_alarm) TableInfo() *TableInfo {
	return History_alarmTableInfo
}
