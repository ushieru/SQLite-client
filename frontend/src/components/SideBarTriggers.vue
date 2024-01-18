<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()

const onClickTrigger = (trigger) => {
    const tab = store.tabs.find(t => t.id == store.currentTabId)
    tab.query = trigger.sql
}
</script>

<template>
    <div class="collapse bg-base-200 rounded-none border-t border-neutral">
        <input type="checkbox" />
        <div class="collapse-title text-xl font-medium">
            Triggers
        </div>
        <div class="collapse-content flex flex-col">
            <button class="btn truncate" v-for="trigger in store.triggers" @click="() => onClickTrigger(trigger)">
                {{ trigger.name }}
            </button>
        </div>
    </div>
</template>