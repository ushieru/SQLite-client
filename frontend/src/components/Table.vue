<script setup>
import { computed } from 'vue';
import { Tab } from '../models/Tab';

const props = defineProps({
    tab: {
        type: Tab,
        required: true,
    }
})

const headers = computed(() => {
    const headers = Object.keys(props.tab.table[0])
    if (headers.find(h => h == 'id'))
        return ['id', ...headers.filter(h => h != 'id')]
    return headers
})
</script>

<template>
    <table v-if="props.tab.table.length" class="table table-zebra">
        <thead>
            <tr>
                <th v-for="header in headers">
                    {{ header }}
                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="row in props.tab.table">
                <th v-for="header in headers">
                    {{ row[header] }}
                </th>
            </tr>
        </tbody>
    </table>
</template>