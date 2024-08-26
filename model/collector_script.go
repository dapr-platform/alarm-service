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


Table: o_collector_script
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] script                                         TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "TfREUYbAtHjcDFohDebuDEEkm",    "script": "EbuNserHCPLfoxtKBtIDTnfXP"}



*/

var (
	Collector_script_FIELD_NAME_id = "id"

	Collector_script_FIELD_NAME_script = "script"
)

// Collector_script struct is a row record of the o_collector_script table in the  database
type Collector_script struct {
	ID     string `json:"id"`     //id
	Script string `json:"script"` //script

}

var Collector_scriptTableInfo = &TableInfo{
	Name: "o_collector_script",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "script",
			Comment:            `script`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Script",
			GoFieldType:        "string",
			JSONFieldName:      "script",
			ProtobufFieldName:  "script",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Collector_script) TableName() string {
	return "o_collector_script"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Collector_script) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Collector_script) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Collector_script) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Collector_script) TableInfo() *TableInfo {
	return Collector_scriptTableInfo
}
