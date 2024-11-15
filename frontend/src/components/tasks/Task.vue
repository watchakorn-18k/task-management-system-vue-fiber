<script setup>
import Card from './Card.vue';
import { ref, onMounted } from 'vue';
import { tasks, getTasksAll, addTask, deleteTask } from '@/services/taskService';
import { randomText } from "@/utils";

const dataFormAddTask = ref({
    name: "",
    details: "",
    status: "ทำอยู่"
});


const removeItem = async (index, taskId) => {
    await deleteTask(taskId);
    tasks.value.splice(index, 1);
};

const openModalAddTask = () => {
    document.getElementById('modalFormAddTask').showModal();
}

const addNewTask = async () => {
    try {
        const task_id = await addTask(dataFormAddTask.value);
        tasks.value.push({
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

onMounted(async () => {
    await getTasksAll();
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
                        autocomplete="off" />
                    <a class="btn btn-sm absolute right-8 btn-neutral" @click="dataFormAddTask.name = ''">X</a>
                </span>
                <label class="label" for="name">รายละเอียด</label>
                <textarea placeholder="รายละเอียด : รายละเอียดของกิจกรรม"
                    class="textarea textarea-bordered textarea-md w-full" v-model="dataFormAddTask.details"
                    autocomplete="off"></textarea>
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

    <div class="flex flex-col items-center">
        <div class="flex flex-row items-center gap-10 my-4 max-w-80">
            <button class="btn btn-primary w-80 text-neutral-100 hover:bg-base-300 hover:border-neutral"
                @click="openModalAddTask">เพิ่มงาน</button>
        </div>

        <div class="flex flex-wrap gap-4 justify-center max-w-[80%]">
            <span v-for="(task, index) in tasks" :key="task.task_id">
                <Card :title="task.name" :details="task.details" v-model:status="task.status"
                    @remove-item="removeItem(index, task.task_id)" />
            </span>
        </div>
    </div>


</template>