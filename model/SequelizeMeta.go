package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details
-------------------------------------


Table: SequelizeMeta
[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "name": "cQQRqiDMKfOSxigHdJdxEIefF"}



*/

// SequelizeMeta struct is a row record of the SequelizeMeta table in the postgres database
type SequelizeMeta struct {
	//[ 0] name                                           VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Name string `gorm:"primary_key;column:name;type:VARCHAR;size:255;"`
}

// SequelizeMetaTableInfo info of SequelizeMeta table
var SequelizeMetaTableInfo = &TableInfo{
	Name: "SequelizeMeta",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SequelizeMeta) TableName() string {
	return "SequelizeMeta"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SequelizeMeta) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SequelizeMeta) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SequelizeMeta) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SequelizeMeta) TableInfo() *TableInfo {
	return SequelizeMetaTableInfo
}
