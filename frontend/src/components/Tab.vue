<script setup>
import Table from './Table.vue'
import { useGlobalStore } from './../stores/GlobalStore'
import { Tab } from '../models/Tab';

const props = defineProps({
    tab: {
        type: Tab,
        required: true,
    },
    isChecked: {
        type: Boolean,
        required: true
    }
})

const store = useGlobalStore()
</script>

<template>
    <input type="radio" name="my_tabs_1" role="tab" class="tab" :aria-label="props.tab.name" :checked="props.isChecked"
        @click="() => store.setCurrentTab(props.tab.id)" />
    <div role="tabpanel" class="tab-content p-10 w-full h-full">
        <textarea class="textarea textarea-bordered w-full" v-model="tab.query" />
        <div class="flex justify-end gap-2">
            <button class="btn" @click="store.execRawQuery">Run</button>
            <button class="btn">Save</button>
        </div>
        <div class="divider"></div>
        <div class="card bg-base-100 shadow-xl">
            <div class="card-body">
                <Table :tab="props.tab" />
            </div>
        </div>
    </div>
</template>