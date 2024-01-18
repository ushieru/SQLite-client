<script setup>
import { Tab } from '../models/Tab';

const props = defineProps({
    tab: {
        type: Tab,
        required: true,
    }
})
</script>

<template>
    <table v-if="props.tab.table?.Headers?.length" class="table table-zebra">
        <thead>
            <tr>
                <th v-for="header in props.tab.table.Headers">
                    {{ header }}
                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="row in props.tab.table.Rows">
                <th v-for="value in row">
                    {{ value }}
                </th>
            </tr>
        </tbody>
    </table>
</template>