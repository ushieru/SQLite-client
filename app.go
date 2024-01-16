package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"pinguino/database"
	"pinguino/models"
	"pinguino/utils"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	auxDB  *sql.DB
	dbName string
}

func NewApp() *App {
	db, _ := database.GetDatabase(database.InMemory)
	return &App{
		auxDB:  db,
		dbName: database.InMemory,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) NewDatabase(name string) {
	a.auxDB.Close()
	db, err := database.GetDatabase(name)
	if err != nil {
		fmt.Println("Error al abrir db")
	}
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

func (a *App) OpenDatabase() string {
	fileName, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		fmt.Println(err.Error())
		nameSlice := strings.Split(a.dbName, "/")
		return nameSlice[len(nameSlice)-1]
	}
	a.NewDatabase(fileName)
	nameSlice := strings.Split(fileName, "/")
	return nameSlice[len(nameSlice)-1]
}

func (a *App) OpenInMemoryDatabase() string {
	a.NewDatabase(database.InMemory)
	return database.InMemory
}

func (a *App) SaveFile(fileName, payload string) bool {
	if _, err := os.Stat("scripts"); os.IsNotExist(err) {
		os.Mkdir("scripts", os.ModePerm)
	}
	cleanBdName := strings.Split(a.dbName, ".")[0]
	if _, err := os.Stat(filepath.Join("scripts", cleanBdName)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join("scripts", cleanBdName), os.ModePerm)
	}
	formatedName := fmt.Sprintf("%s.sql", fileName)
	file, err := os.Create(filepath.Join("scripts", cleanBdName, formatedName))
	if err != nil {
		return false
	}
	file.Write([]byte(payload))
	file.Close()
	return true
}
