import { defineStore } from 'pinia'
import { Tab } from '../models/Tab'

export const useGlobalStore = defineStore('global', {
    state: () => ({
        currentTabId: 1,
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
        getTables() {
            window.go.main.App.GetTables()
                .then(tables => this.tables = tables)
        },
        getViews() {
            window.go.main.App.GetViews()
                .then(views => this.views = views)
        },
        selectTable(tableName) {
            window.go.main.App.SelectTable(tableName)
                .then(result => {
                    const tab = this.tabs.find(t => t.id == this.currentTabId)
                    tab.table = result
                    tab.query = `SELECT * FROM ${tableName} LIMIT 100`
                })
        },
        execRawQuery() {
            const tab = this.tabs.find(t => t.id == this.currentTabId)
            window.go.main.App.ExecRawQuery(tab.query)
                .then(result => {
                    if (!result.length) return
                    tab.table = result
                })
        }
    }
})