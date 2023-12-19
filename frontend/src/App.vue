<script setup>
import SideBar from './components/SideBar.vue'
import Tab from './components/Tab.vue'
import { useGlobalStore } from './stores/GlobalStore'
const store = useGlobalStore()
store.getTables()
store.getViews()
</script>

<template>
  <main class="flex h-screen w-screen">
    <div class="w-[200px] flex flex-col">
      <SideBar />
    </div>
    <div class="divider divider-horizontal m-0" />
    <div class="grow h-full p-5 overflow-auto">
      <div role="tablist" class="tabs tabs-bordered w-full">
        <Tab v-for="tab in store.tabs" :tab="tab" :isChecked="tab.id == store.currentTabId" />
        <input type="radio" name="my_tabs_1" role="tab" class="tab" aria-label="New tab" @click="store.newTab" />
      </div>
    </div>
  </main>
</template>
