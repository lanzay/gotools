package sqlTools

import (
	"database/sql"
)
//Возвращает map из 1й строки результата запроса
func MapFromRows(rows *sql.Rows) (map[string]interface{}, error) {
	
	col, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	
	cols := make([]interface{},len(col))
	colsPtr := make([]interface{},len(col))
	for i, _ := range cols {
		colsPtr[i] = &cols[i]
	}
	
	for rows.Next() {
		err = rows.Scan(colsPtr...)
		if err != nil {
			return nil, err
		}
		break
	}
	
	m := make(map[string]interface{},len(cols))
	for i, colName := range col {
		val := colsPtr[i].(*interface{})
		m[colName] = *val
	}
	
	return m, nil
}
