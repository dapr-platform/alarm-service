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


Table: f_event
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] dn                                             VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] event_title                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] event_text                                     TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] event_time                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] object_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 7] event_extra                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] level                                          INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]


JSON Sample
-------------------------------------
{    "id": "OepHJuAwvVGOEyTFGNRQvogld",    "dn": "OlZxMSYcwRMMjEKtLkUdVtgKc",    "event_title": "OjgYMVuUEQxsctWJueEPJfcNk",    "event_text": "ANFKHobdWpFDyIPlcIilwjxhH",    "event_time": 58,    "object_id": "cZRLEhndbdBXRgWobnfZnEQok",    "status": 23,    "event_extra": "iXNJbZjorNTGVsUgiiHMpLaNE",    "level": 25}



*/

var (
	Event_FIELD_NAME_id = "id"

	Event_FIELD_NAME_dn = "dn"

	Event_FIELD_NAME_event_title = "event_title"

	Event_FIELD_NAME_event_text = "event_text"

	Event_FIELD_NAME_event_time = "event_time"

	Event_FIELD_NAME_object_id = "object_id"

	Event_FIELD_NAME_status = "status"

	Event_FIELD_NAME_event_extra = "event_extra"

	Event_FIELD_NAME_level = "level"
)

// Event struct is a row record of the f_event table in the  database
type Event struct {
	ID         string           `json:"id"`          //id
	Dn         string           `json:"dn"`          //dn
	EventTitle string           `json:"event_title"` //event_title
	EventText  string           `json:"event_text"`  //event_text
	EventTime  common.LocalTime `json:"event_time"`  //event_time
	ObjectID   string           `json:"object_id"`   //object_id
	Status     int32            `json:"status"`      //status
	EventExtra string           `json:"event_extra"` //event_extra
	Level      int32            `json:"level"`       //level

}

var EventTableInfo = &TableInfo{
	Name: "f_event",
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
			Name:               "event_title",
			Comment:            `event_title`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "EventTitle",
			GoFieldType:        "string",
			JSONFieldName:      "event_title",
			ProtobufFieldName:  "event_title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "event_text",
			Comment:            `event_text`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "EventText",
			GoFieldType:        "string",
			JSONFieldName:      "event_text",
			ProtobufFieldName:  "event_text",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "event_time",
			Comment:            `event_time`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "object_id",
			Comment:            `object_id`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `status`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "event_extra",
			Comment:            `event_extra`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "EventExtra",
			GoFieldType:        "string",
			JSONFieldName:      "event_extra",
			ProtobufFieldName:  "event_extra",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "level",
			Comment:            `level`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Event) TableName() string {
	return "f_event"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Event) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Event) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Event) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Event) TableInfo() *TableInfo {
	return EventTableInfo
}
