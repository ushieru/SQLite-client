<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()
</script>

<template>
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
</template>