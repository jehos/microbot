package microbot

import (
	"fmt"

	"github.com/elvinchan/microbot/db"
)

func PingDB() {
	for _, d := range dialects {
		d.DB().Ping()
	}
}

type TableInfo struct {
	DBType db.DBType  `json:"dbType"`
	Tables []db.Table `json:"tables"`
}

func GetTableInfo() ([]TableInfo, error) {
	var tableInfos []TableInfo
	for _, d := range dialects {
		tables, err := d.GetTables()
		if err != nil {
			return nil, err
		}

		for i := range tables {
			cols, err := d.GetColumns(tables[i].Name)
			if err != nil {
				return nil, err
			}

			tables[i].Columns = cols
			indexes, err := d.GetIndexes(tables[i].Name)
			if err != nil {
				return nil, err
			}
			tables[i].Indexes = indexes

			for _, index := range indexes {
				for _, name := range index.Cols {
					if col := tables[i].GetColumn(name); col != nil {
						col.Indexes[index.Name] = index.Type
					} else {
						return nil, fmt.Errorf("Unknown col %s in index %v of table %v", name, index.Name, tables[i].Name)
					}
				}
			}
		}
		tableInfos = append(tableInfos, TableInfo{
			Tables: tables,
			DBType: d.DBType(),
		})
	}

	return tableInfos, nil
}
