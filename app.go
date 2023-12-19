package main

import (
	"context"
	"database/sql"
	"fmt"
	"pinguino/database"
	"pinguino/utils"
)

type App struct {
	ctx   context.Context
	auxDB *sql.DB
}

func NewApp() *App {
	db, _ := database.GetDatabase(database.InMemory)
	return &App{
		auxDB: db,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetTables() []map[string]interface{} {
	rows, err := a.auxDB.Query("SELECT name FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%'")
	if err != nil {
		return make([]map[string]interface{}, 0)
	}
	mapRows, _ := utils.SQLRowsToMaps(rows)
	return mapRows
}

func (a *App) GetViews() []map[string]interface{} {
	rows, err := a.auxDB.Query("SELECT name FROM sqlite_master WHERE type = 'view' AND name NOT LIKE 'sqlite_%'")
	if err != nil {
		return make([]map[string]interface{}, 0)
	}
	mapRows, _ := utils.SQLRowsToMaps(rows)
	return mapRows
}

func (a *App) SelectTable(tableName string) []map[string]interface{} {
	rows, err := a.auxDB.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 100", tableName))
	if err != nil {
		return make([]map[string]interface{}, 0)
	}
	mapRows, _ := utils.SQLRowsToMaps(rows)
	return mapRows
}

func (a *App) ExecRawQuery(query string) []map[string]interface{} {
	rows, err := a.auxDB.Query(query)
	if err != nil {
		return make([]map[string]interface{}, 0)
	}
	mapRows, _ := utils.SQLRowsToMaps(rows)
	return mapRows
}
