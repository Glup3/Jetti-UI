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


Table: Team
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] teamName                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] playerId1                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] playerId2                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] playerId3                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] playerId4                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] playerId5                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 8] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []


JSON Sample
-------------------------------------
{    "updated_at": "2059-03-25T09:34:01.70715842+01:00",    "player_id_1": 31,    "player_id_2": 99,    "player_id_4": 74,    "player_id_5": 45,    "id": 52,    "team_name": "PMvMjpOvkwvGvISKbQreJFDqH",    "player_id_3": 62,    "created_at": "2093-09-06T17:56:47.90325234+02:00"}



*/

// Team struct is a row record of the Team table in the postgres database
type Team struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;"`
	//[ 1] teamName                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	TeamName null.String `gorm:"column:teamName;type:VARCHAR;size:255;"`
	//[ 2] playerId1                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID1 null.Int `gorm:"column:playerId1;type:INT4;"`
	//[ 3] playerId2                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID2 null.Int `gorm:"column:playerId2;type:INT4;"`
	//[ 4] playerId3                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID3 null.Int `gorm:"column:playerId3;type:INT4;"`
	//[ 5] playerId4                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID4 null.Int `gorm:"column:playerId4;type:INT4;"`
	//[ 6] playerId5                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	PlayerID5 null.Int `gorm:"column:playerId5;type:INT4;"`
	//[ 7] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:createdAt;type:TIMESTAMPTZ;"`
	//[ 8] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updatedAt;type:TIMESTAMPTZ;"`

	Player1 PlayerH `gorm:"foreignKey:PlayerID1"`
	Player2 PlayerH `gorm:"foreignKey:PlayerID2"`
	Player3 PlayerH `gorm:"foreignKey:PlayerID3"`
	Player4 PlayerH `gorm:"foreignKey:PlayerID4"`
	Player5 PlayerH `gorm:"foreignKey:PlayerID5"`
}

// TeamTableInfo info of Team table
var TeamTableInfo = &TableInfo{
	Name: "Team",
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
			Name:               "teamName",
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
			GoFieldName:        "TeamName",
			GoFieldType:        "null.String",
			JSONFieldName:      "team_name",
			ProtobufFieldName:  "team_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "playerId1",
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
			GoFieldName:        "PlayerID1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id_1",
			ProtobufFieldName:  "player_id_1",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "playerId2",
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
			GoFieldName:        "PlayerID2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id_2",
			ProtobufFieldName:  "player_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "playerId3",
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
			GoFieldName:        "PlayerID3",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id_3",
			ProtobufFieldName:  "player_id_3",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "playerId4",
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
			GoFieldName:        "PlayerID4",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id_4",
			ProtobufFieldName:  "player_id_4",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "playerId5",
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
			GoFieldName:        "PlayerID5",
			GoFieldType:        "null.Int",
			JSONFieldName:      "player_id_5",
			ProtobufFieldName:  "player_id_5",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Team) TableName() string {
	return "Team"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Team) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Team) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Team) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Team) TableInfo() *TableInfo {
	return TeamTableInfo
}
