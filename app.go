package main

import (
	"context"
	"database/sql"
	"fmt"
	"pinguino/database"
	"pinguino/models"
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

func (a *App) NewDatabase(name string) {
	a.auxDB.Close()
	db, _ := database.GetDatabase(name)
	a.auxDB = db
}

func (a *App) GetTables() []string {
	rows, err := a.auxDB.Query("SELECT name FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%'")
	if err != nil {
		return make([]string, 0)
	}
	r, _ := utils.SQLRowsToSlice(rows)
	tables := make([]string, 0)
	for _, row := range r.Rows {
		if val, ok := row[0].(string); ok {
			tables = append(tables, val)
		}
	}
	return tables
}

func (a *App) GetViews() []string {
	rows, err := a.auxDB.Query("SELECT name FROM sqlite_master WHERE type = 'view' AND name NOT LIKE 'sqlite_%'")
	if err != nil {
		return make([]string, 0)
	}
	r, _ := utils.SQLRowsToSlice(rows)
	tables := make([]string, 0)
	for _, row := range r.Rows {
		if val, ok := row[0].(string); ok {
			tables = append(tables, val)
		}
	}
	return tables
}

func (a *App) SelectTable(tableName string) (*models.SQLResult, error) {
	rows, err := a.auxDB.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 100", tableName))
	if err != nil {
		return nil, err
	}
	r, _ := utils.SQLRowsToSlice(rows)
	return r, nil
}

func (a *App) ExecRawQuery(query string) (*models.SQLResult, error) {
	rows, err := a.auxDB.Query(query)
	if err != nil {
		return nil, err
	}
	r, _ := utils.SQLRowsToSlice(rows)
	fmt.Println(r.Rows)
	fmt.Println(r.Headers)
	return r, nil
}
