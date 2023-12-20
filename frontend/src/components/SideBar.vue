<script setup>
import SideBarTables from './SideBarTables.vue';
import SideBarViews from './SideBarViews.vue';
import SideBarTriggers from './SideBarTriggers.vue';
import { useGlobalStore } from './../stores/GlobalStore'

const store = useGlobalStore()

const showNewDatabaseModal = () =>
    document.getElementById('new_database_modal').showModal()

const closeNewDatabaseModal = () =>
    document.getElementById('new_database_modal').close()

const onSubmitNewDatabaseModal = (e) => {
    e.preventDefault()
    const databaseName = e.target.database_name.value
    store.newDatabase(databaseName)
    e.target.database_name.value = ""
    document.getElementById('new_database_modal').close()
}

</script>

<template>
    <div class="flex justify-between items-center p-1">
        <div class="badge badge-ghost py-6 rounded-lg w-full px-1 text-ellipsis">{{ store.databaseName }}</div>
        <div class="dropdown">
            <div tabindex="0" role="button" class="btn m-1">▼</div>
            <ul tabindex="0" class="dropdown-content z-10 menu p-2 shadow bg-base-100 rounded-box w-52">
                <li><a>New in-memory</a></li>
                <li @click="showNewDatabaseModal"><a>New database</a></li>
                <li><a>Open database</a></li>
                <div class="divider m-0" />
                <li><a>Close database</a></li>
            </ul>
        </div>
    </div>
    <div class="divider my-0"></div>
    <div class="grow overflow-auto pl-2">
        <SideBarTables />
        <SideBarViews />
        <SideBarTriggers />
    </div>
    <dialog id="new_database_modal" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">New database</h3>
            <form @submit="onSubmitNewDatabaseModal">
                <label for="database_name">Name</label>
                <br>
                <input type="text" name="database_name" class="input input-bordered w-full max-w-xs" />
                <div class="modal-action">
                    <button class="btn" type="button" @click="closeNewDatabaseModal">Close</button>
                    <button class="btn btn-primary" type="submit">New</button>
                </div>
            </form>
        </div>
        <form method="dialog" class="modal-backdrop">
            <button>close</button>
        </form>
    </dialog>
</template>