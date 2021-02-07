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


Table: Match
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] matchResult                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] screenshotPath                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] teamId1                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] teamId2                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 6] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []


JSON Sample
-------------------------------------
{    "updated_at": "2024-01-16T19:27:16.100505466+01:00",    "id": 42,    "match_result": 81,    "screenshot_path": "dsvUILexrtXbJKjHteSmrhSwO",    "team_id_1": 99,    "team_id_2": 46,    "created_at": "2028-06-22T23:33:09.650896341+02:00"}



*/

// Match struct is a row record of the Match table in the postgres database
type Match struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;"`
	//[ 1] matchResult                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	MatchResult null.Int `gorm:"column:matchResult;type:INT4;"`
	//[ 2] screenshotPath                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ScreenshotPath null.String `gorm:"column:screenshotPath;type:VARCHAR;size:255;"`
	//[ 3] teamId1                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TeamID1 null.Int `gorm:"column:teamId1;type:INT4;"`
	//[ 4] teamId2                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TeamID2 null.Int `gorm:"column:teamId2;type:INT4;"`
	//[ 5] createdAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:createdAt;type:TIMESTAMPTZ;"`
	//[ 6] updatedAt                                      TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updatedAt;type:TIMESTAMPTZ;"`

	Team1 Team `gorm:"foreignKey:TeamID1"`

	Team2 Team `gorm:"foreignKey:TeamID2"`
}

// MatchTableInfo info of Match table
var MatchTableInfo = &TableInfo{
	Name: "Match",
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
			Name:               "matchResult",
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
			GoFieldName:        "MatchResult",
			GoFieldType:        "null.Int",
			JSONFieldName:      "match_result",
			ProtobufFieldName:  "match_result",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "screenshotPath",
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
			GoFieldName:        "ScreenshotPath",
			GoFieldType:        "null.String",
			JSONFieldName:      "screenshot_path",
			ProtobufFieldName:  "screenshot_path",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "teamId1",
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
			GoFieldName:        "TeamID1",
			GoFieldType:        "null.Int",
			JSONFieldName:      "team_id_1",
			ProtobufFieldName:  "team_id_1",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "teamId2",
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
			GoFieldName:        "TeamID2",
			GoFieldType:        "null.Int",
			JSONFieldName:      "team_id_2",
			ProtobufFieldName:  "team_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Match) TableName() string {
	return "Match"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Match) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Match) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Match) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Match) TableInfo() *TableInfo {
	return MatchTableInfo
}
