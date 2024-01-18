<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()

const onClickAutoQuery = (tableName, action) => {
    const tab = store.tabs.find(t => t.id == store.currentTabId)
    switch (action) {
        case 'C':
            tab.query = `INSER INTO \`${tableName}\` ()\nVALUES()`
            break;
        case 'R':
            tab.query = `SELECT * FROM \`${tableName}\`\nWHERE ? = ?`
            break;
        case 'U':
            tab.query = `UPDATE \`${tableName}\` SET ? = ?\nWHERE ? = ?`
            break;
        case 'D':
            tab.query = `DELETE FROM \`${tableName}\`\nWHERE ? = ?`
            break;
        case 'T':
            tab.query = `CREATE TRIGGER [IF NOT EXISTS] trigger_name\n[BEFORE|AFTER|INSTEAD OF] [INSERT|UPDATE|DELETE]\nON ${tableName}\n[WHEN condition]\nBEGIN\n\tstatements;\nEND;`
            break;
    }
}
</script>

<template>
    <div class="collapse bg-base-200 rounded-none">
        <input type="checkbox" />
        <div class="collapse-title text-xl font-medium">
            Tables
        </div>
        <div class="collapse-content flex flex-col">
            <div v-for="name in store.tables" class="flex">
                <button class="btn grow" @click="() => store.selectTable(name)">
                    {{ name }}
                </button>
                <div class="dropdown">
                    <div tabindex="0" role="button" class="btn">▼</div>
                    <ul tabindex="0" style="position: fixed;"
                        class="dropdown-content z-10 menu p-2 shadow bg-base-100 rounded-box w-52">
                        <li><a @click="() => onClickAutoQuery(name, 'C')">Create</a></li>
                        <li><a @click="() => onClickAutoQuery(name, 'R')">Read</a></li>
                        <li><a @click="() => onClickAutoQuery(name, 'U')">Update</a></li>
                        <li><a @click="() => onClickAutoQuery(name, 'D')">Delete</a></li>
                        <li><a @click="() => onClickAutoQuery(name, 'T')">Trigger</a></li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>