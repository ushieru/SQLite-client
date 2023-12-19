<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()
</script>

<template>
    <div class="flex justify-between items-center p-1">
        <div class="badge badge-ghost py-6 rounded-lg w-full px-1 text-ellipsis">:memory:</div>
        <div class="dropdown">
            <div tabindex="0" role="button" class="btn m-1">▼</div>
            <ul tabindex="0" class="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52">
                <li><a>New in-memory</a></li>
                <li><a>Open database</a></li>
                <div class="divider m-0" />
                <li><a>Close database</a></li>
            </ul>
        </div>
    </div>
    <div class="divider my-0"></div>
    <div class="grow overflow-auto pl-2">
        <div class="collapse bg-base-200 rounded-none">
            <input type="checkbox" />
            <div class="collapse-title text-xl font-medium">
                Tables
            </div>
            <div class="collapse-content flex flex-col">
                <button class="btn" v-for="{ name } in store.tables" @click="() => store.selectTable(name)">
                    {{ name }}
                </button>
            </div>
        </div>

        <div class="collapse bg-base-200 rounded-none border-t border-neutral">
            <input type="checkbox" />
            <div class="collapse-title text-xl font-medium">
                Views
            </div>
            <div class="collapse-content flex flex-col">
                <button class="btn" v-for="{ name } in store.views" @click="() => store.selectTable(name)">
                    {{ name }}
                </button>
            </div>
        </div>

        <div class="collapse bg-base-200 rounded-none border-t border-neutral">
            <input type="checkbox" />
            <div class="collapse-title text-xl font-medium">
                Triggers
            </div>
            <div class="collapse-content flex flex-col">
                <button class="btn">Trigger 1</button>
                <button class="btn">Trigger 2</button>
                <button class="btn">Trigger 3</button>
            </div>
        </div>
    </div>
</template>