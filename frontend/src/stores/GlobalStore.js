import { defineStore } from 'pinia'
import { Tab } from '../models/Tab'
import {
    ExecRawQuery, GetTables, GetViews, SelectTable, NewDatabase, OpenDatabase, OpenInMemoryDatabase,
} from '../../wailsjs/go/main/App'

export const useGlobalStore = defineStore('global', {
    state: () => ({
        currentTabId: 1,
        databaseName: ":memory:",
        tabs: [new Tab("Tab 1", "", 1)],
        tables: [],
        views: [],
        _internalTabConter: 2,
    }),
    actions: {
        setCurrentTab(tabId) {
            this.currentTabId = tabId
        },
        newTab() {
            const name = `Tab ${this._internalTabConter}`
            const tab = new Tab(name, "")
            this._internalTabConter++
            this.currentTabId = tab.id
            this.tabs.push(tab)
        },
        closeTab(tabId) {
            if (this.tabs.length == 1) return
            this.tabs = this.tabs.filter(tab => tab.id != tabId)
            this.currentTabId = this.tabs[0].id
        },
        getTables() {
            GetTables()
                .then(tables => this.tables = tables)
        },
        getViews() {
            GetViews()
                .then(views => this.views = views)
        },
        selectTable(tableName) {
            SelectTable(tableName)
                .then(result => {
                    const tab = this.tabs.find(t => t.id == this.currentTabId)
                    tab.table = result
                    console.log(result)
                    tab.query = `SELECT * FROM ${tableName} LIMIT 100`
                })
        },
        execRawQuery(tab) {
            ExecRawQuery(tab.query)
                .then(result => {
                    tab.table = result
                }).then(_ => {
                    if (tab.query.toLowerCase().includes("create")) {
                        this.getTables()
                        this.getViews()
                    }
                })
        },
        newDatabase(databaseName) {
            if (!databaseName.includes(".")) {
                databaseName = `${databaseName}.db`
            }
            NewDatabase(databaseName).then(_ => {
                this.databaseName = databaseName
                this.getTables()
                this.getViews()
            })
        },
        openInMemoryDatabase() {
            OpenInMemoryDatabase().then(dbName => {
                this.databaseName = dbName
                this.getTables()
                this.getViews()
            })
        },
        openDatabase() {
            OpenDatabase().then(dbName => {
                this.databaseName = dbName
                this.getTables()
                this.getViews()
            })
        },
    }
})