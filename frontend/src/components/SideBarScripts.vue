<script setup>
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()
</script>

<template>
    <div class="collapse bg-base-200 rounded-none border-t border-neutral">
        <input type="checkbox" />
        <div class="collapse-title text-xl font-medium">
            Scripts
        </div>
        <div class="collapse-content flex flex-col">
            <button class="btn" v-for="name in store.scripts" @click="() => store.readFile(name)">
                {{ name }}
            </button>
        </div>
    </div>
</template>