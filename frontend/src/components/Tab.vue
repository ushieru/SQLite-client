<script setup>
import { ref } from 'vue'
import Table from './Table.vue'
import { useGlobalStore } from './../stores/GlobalStore'
import { Tab } from '../models/Tab';
import { Codemirror } from 'vue-codemirror'
import { sql, SQLite } from '@codemirror/lang-sql'
import { oneDark } from '@codemirror/theme-one-dark'

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

const name = ref(props.tab.name)

const filename = ref("")

const onClickTab = () => {
    if (props.isChecked) {
        document.getElementById(props.tab.name).showModal()
    }
    store.setCurrentTab(props.tab.id)
}

const onClickSave = () => {
    document.getElementById(props.tab.name + 'save').showModal()
}

const saveFile = () => {
    document.getElementById(props.tab.name + 'save').close()
    store.saveFile(filename, props.tab.query)
}

const onMouseHover = () => {
    if (store.tabs.length == 1) return
    if (props.isChecked) name.value = 'Close'
}

const extentions = [sql({ dialect: SQLite }), oneDark]
</script>

<template>
    <input type="radio" name="my_tabs_1" role="tab" class="tab" :aria-label="name" :checked="props.isChecked"
        @click="onClickTab" @mouseover="onMouseHover" @mouseleave="name = props.tab.name" />
    <div role="tabpanel" class="tab-content p-10 w-full h-full">
        <Codemirror :extensions="extentions" v-model="tab.query" :style="{ height: '150px' }" />
        <div class="flex justify-end gap-2 mt-1">
            <button class="btn" @click="() => store.execRawQuery(props.tab)">Run</button>
            <button class="btn" @click="onClickSave">Save</button>
        </div>
        <div class="divider"></div>
        <div class="card bg-base-100 shadow-xl">
            <div class="card-body">
                <Table :tab="props.tab" />
            </div>
        </div>
    </div>
    <dialog :id="props.tab.name" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">{{ props.tab.name }}</h3>
            <p class="py-4">Seguro que quieres cerrar esta pestaña?</p>
            <div class="modal-action">
                <form method="dialog">
                    <button class="btn">No</button>
                </form>
                <button class="btn btn-primary" @click="saveFile">Si</button>
            </div>
        </div>
        <form method="dialog" class="modal-backdrop">
            <button>close</button>
        </form>
    </dialog>
    <dialog :id="props.tab.name + 'save'" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">Guardar archivo</h3>
            <p class="py-4">Seguro que quieres guardar este script?</p>
            <code class="max-h-24">{{ props.tab.query }}</code>
            <h3 class="font-bold text-lg py-2">Nombre del archivo</h3>
            <input type="text" v-model="filename" class="input input-bordered w-full max-w-xs">
            <div class="modal-action">
                <form method="dialog">
                    <button class="btn">No</button>
                </form>
                <button class="btn btn-primary" @click="() => store.saveFile(filename, props.tab.query)">Si</button>
            </div>
        </div>
        <form method="dialog" class="modal-backdrop">
            <button>close</button>
        </form>
    </dialog>
</template>