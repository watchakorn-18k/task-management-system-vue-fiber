<script setup>
import Card from './Card.vue';
import { ref, onMounted } from 'vue';
import { getTasksAll, addTask, deleteTask } from '@/services/taskService';
import { randomText } from "@/utils";

const dataFormAddTask = ref({
    name: "",
    details: "",
    status: "ทำอยู่"
});

const tasksSetup = ref([]);
const tasksAll = ref([]);


const removeItem = async (index, taskId) => {
    await deleteTask(taskId);
    tasksAll.value.splice(index, 1);
};

const openModalAddTask = () => {
    document.getElementById('modalFormAddTask').showModal();
}

const addNewTask = async () => {
    try {
        const task_id = await addTask(dataFormAddTask.value);
        tasksAll.value.push({
            task_id: task_id,
            name: dataFormAddTask.value.name,
            details: dataFormAddTask.value.details,
            status: dataFormAddTask.value.status
        });
        document.getElementById('modalFormAddTask').close();
        dataFormAddTask.value = { status: "ทำอยู่" };
    } catch (error) {
        console.log(error);

    }

}

const searchTask = (e) => {
    const value = e.target.value;
    if (value === "") {
        tasksAll.value = tasksSetup.value
        return;
    }
    tasksAll.value = tasksAll.value.filter(task => task.name.toLowerCase().includes(value.toLowerCase()));


}

const filterStatus = (e) => {
    const value = e.target.value;
    tasksAll.value = tasksSetup.value
    if (value === "ทั้งหมด") {
        tasksAll.value = tasksSetup.value
        return;
    }
    tasksAll.value = tasksAll.value.filter(task => task.status === value);

}

onMounted(async () => {
    tasksSetup.value = await getTasksAll();
    tasksAll.value = tasksSetup.value
});

</script>

<template>

    <dialog id="modalFormAddTask" class="modal">
        <div class="modal-box">
            <form method="dialog">
                <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
            </form>
            <h3 class="text-xl font-bold text-center">เพิ่มงาน</h3>
            <form v-on:submit.prevent="addNewTask" class="w-full">
                <label class="label" for="name">ชื่อ</label>
                <span class="flex flex-row items-center">
                    <input type="text" placeholder="ชื่อ : กิจกรรมเล่นเกม" id="name"
                        class="input input-bordered input-md w-full" required v-model="dataFormAddTask.name"
                        autocomplete="off" :maxlength="20" />
                    <a class="btn btn-sm absolute right-8 btn-neutral" @click="dataFormAddTask.name = ''">X</a>
                </span>
                <label class="label" for="name">รายละเอียด</label>
                <textarea placeholder="รายละเอียด : รายละเอียดของกิจกรรม"
                    class="textarea textarea-bordered textarea-md w-full" v-model="dataFormAddTask.details"
                    autocomplete="off" :maxlength="60"></textarea>
                <label class="label" for="name">สถานะ</label>
                <select class="select select-bordered w-full" v-model="dataFormAddTask.status">
                    <option selected>ทำอยู่</option>
                    <option>ทำเสร็จแล้ว</option>
                </select>
                <div class="modal-action">
                    <button class="btn btn-primary text-neutral-100 w-full ">บันทึก</button>
                </div>
            </form>
        </div>
    </dialog>

    <div class="flex flex-col items-center m-8">
        <div class="flex flex-col md:flex-row items-center justify-between gap-2 my-4 w-full md:max-w-[80%]">
            <label class="input input-bordered flex w-full md:w-full items-center gap-2">
                <input type="text" class="grow" placeholder="ค้นหา" @input="searchTask" />
                <span class="pi pi-search"></span>
            </label>
            <select class="select select-bordered join-item" @change="filterStatus">
                <option>ทั้งหมด</option>
                <option>ทำอยู่</option>
                <option>ทำเสร็จแล้ว</option>
            </select>
            <button class="btn btn-primary w-full md:w-auto text-neutral-100 hover:bg-base-200 hover:border-neutral"
                @click="openModalAddTask"><span class="pi pi-plus"></span></button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 justify-center w-full md:max-w-[80%]">
            <span v-for="(task, index) in tasksAll" :key="task.task_id">
                <Card :title="task.name" :details="task.details" v-model:status="task.status"
                    @remove-item="removeItem(index, task.task_id)" />
            </span>
        </div>
    </div>


</template>