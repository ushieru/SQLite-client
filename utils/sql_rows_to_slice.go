package utils

import (
	"database/sql"
	"pinguino/models"
)

func SQLRowsToSlice(rows *sql.Rows) (*models.SQLResult, error) {
	cols, _ := rows.Columns()
	results := make([][]interface{}, 0)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		m := make([]interface{}, 0)
		for i := range cols {
			val := columnPointers[i].(*interface{})
			m = append(m, *val)
		}
		results = append(results, m)
	}
	result := models.SQLResult{
		Headers: cols,
		Rows:    results,
	}
	return &result, nil
}
