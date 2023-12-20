<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()
</script>

<template>
    <div class="collapse bg-base-200 rounded-none">
        <input type="checkbox" />
        <div class="collapse-title text-xl font-medium">
            Tables
        </div>
        <div class="collapse-content flex flex-col">
            <div v-for="{ name } in store.tables" class="flex">
                <button class="btn grow" @click="() => store.selectTable(name)">
                    {{ name }}
                </button>
                <div v-if="false" class="dropdown">
                    <div tabindex="0" role="button" class="btn">▼</div>
                    <ul tabindex="0" style="position: fixed;" class="dropdown-content z-10 menu p-2 shadow bg-base-100 rounded-box w-52">
                        <li><a>Create</a></li>
                        <li><a>Update</a></li>
                        <li><a>Delete</a></li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>