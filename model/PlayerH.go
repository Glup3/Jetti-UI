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


Table: PlayerH
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] playerId                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] skillLevel                                     FLOAT8               null: true   primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []
[ 3] userTag                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 5] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []


JSON Sample
-------------------------------------
{    "created_at": "2193-10-18T09:38:01.182489544+02:00",    "updated_at": "2284-09-21T20:41:12.95189906+02:00",    "id": 44,    "player_id": 33,    "skill_level": 0.021775587,    "user_tag": "xMLTouovDFODxhrbiqcubNFyv"}



*/

// PlayerH struct is a row record of the PlayerH table in the postgres database
type PlayerH struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;"`
	//[ 1] playerId                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID null.Int `gorm:"column:playerId;type:INT4;"`
	//[ 2] skillLevel                                     FLOAT8               null: true   primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []
	SkillLevel null.Float `gorm:"column:skillLevel;type:FLOAT8;"`
	//[ 3] userTag                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	UserTag null.String `gorm:"column:userTag;type:VARCHAR;size:255;"`
	//[ 4] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:createdAt;type:TIMESTAMPTZ;"`
	//[ 5] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updatedAt;type:TIMESTAMPTZ;"`

	Player Players `gorm:"foreignKey:PlayerID"`
}

// PlayerHTableInfo info of PlayerH table
var PlayerHTableInfo = &TableInfo{
	Name: "PlayerH",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "playerId",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "PlayerID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id",
			ProtobufFieldName:  "player_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "skillLevel",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "FLOAT8",
			DatabaseTypePretty: "FLOAT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT8",
			ColumnLength:       -1,
			GoFieldName:        "SkillLevel",
			GoFieldType:        "null.Float",
			JSONFieldName:      "skill_level",
			ProtobufFieldName:  "skill_level",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "userTag",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "UserTag",
			GoFieldType:        "null.String",
			JSONFieldName:      "user_tag",
			ProtobufFieldName:  "user_tag",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "createdAt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "updatedAt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PlayerH) TableName() string {
	return "PlayerH"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PlayerH) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PlayerH) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PlayerH) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PlayerH) TableInfo() *TableInfo {
	return PlayerHTableInfo
}
