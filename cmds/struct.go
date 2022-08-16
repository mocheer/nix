package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/ts"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Struct : nix struct schemeName
var Struct = &cli.Command{
	Name:  "struct",
	Usage: "从数据库中读取库表模型转成struct",
	Action: func(c *cli.Context) error {
		db, err := openDB()
		if err != nil {
			return err
		}
		//

		query := db.Raw(`SELECT 
		tb.tablename as tablename,
		a.attname AS columnname,
		t.typname AS type
		FROM
		pg_class as c,
		pg_attribute as a, 
		pg_type as t,
		(select tablename from pg_tables where schemaname = @schema) as tb
		WHERE  a.attnum > 0 
		and a.attrelid = c.oid
		and a.atttypid = t.oid
		and c.relname = tb.tablename 
		order by tablename`, map[string]any{"schema": "pipal"})
		result := scanIntoMap(query)
		fmt.Println(result)
		return nil
	},
}

func openDB() (db *gorm.DB, err error) {
	data := ds_json.ReadGJSON("./assets/config/app.json")
	if err != nil {
		return nil, err
	}
	db, err = gorm.Open(postgres.Open(data.Get("DSN").String()), &gorm.Config{
		AllowGlobalUpdate: false,
		Logger:            logger.Default.LogMode(logger.Silent),
	})
	return
}

func scanIntoMap(query *gorm.DB) []ts.Map[any] {
	var result []ts.Map[any]
	rows, err := query.Rows()
	fmt.Println(err)
	if err == nil {
		for rows.Next() {
			rowData := map[string]any{}
			err = query.ScanRows(rows, &rowData)
			//
			if err == nil {
				columnTypes, _ := rows.ColumnTypes()
				for _, columnType := range columnTypes {
					// gorm-postgresql 底层的 DataType 并没有识别JSON，然后转成[]byte，而是用的string
					if columnType.DatabaseTypeName() == "JSON" {
						name := columnType.Name()
						if fn.IsString(rowData[name]) {
							var data any
							err := json.Unmarshal([]byte(rowData[name].(string)), &data)
							if err == nil {
								rowData[name] = data
							}
						}
					}
				}
			}
			//
			result = append(result, rowData)
		}
	}
	return result
}
