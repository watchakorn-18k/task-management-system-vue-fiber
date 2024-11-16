<script setup>
import { ref } from 'vue';
defineProps({
    title: {
        type: String,
    },
    details: {
        type: String
    },
    status: {
        type: String
    }

})
const emit = defineEmits(['update:status', 'remove-item', 'editTaskEmit', 'updateStatusEmit']);

const handleStatusChange = (event) => {
    const newStatus = event.target.value;
    emit('update:status', newStatus);
    emit('updateStatusEmit', newStatus);
}
const deleteCard = () => {
    emit('remove-item');
}

const handleEditTask = () => {
    emit('editTaskEmit');
}

</script>
<template>
    <div class="card bg-neutral-100 w-full h-60 shadow-xl">
        <div class="card-body">
            <div class="card-actions justify-end">

                <p class="font-bold text-lg text-base-100">{{ title }}</p>
                <button class="btn btn-square btn-sm hover:bg-error hover:text-neutral-100 hover:border-error"
                    @click="deleteCard">
                    X
                </button>
            </div>
            <p class="text-base-200 text-sm mt-4">{{ details }}</p>

        </div>

        <div class="flex justify-between items-center m-8">

            <div class="flex flex-row items-center gap-4">
                <p>สถานะ</p>
                <select class="select badge badge-outline h-10 cursor-pointer focus:outline-none" :class="{
                    'badge-primary': status === 'ทำอยู่',
                    'badge-success': status === 'ทำเสร็จแล้ว',
                }" @change="handleStatusChange">

                    <option class="text-base-300 selection:bg-pink-300" :selected="status === 'ทำอยู่'">ทำอยู่</option>
                    <option class="text-base-300" :selected="status === 'ทำเสร็จแล้ว'">ทำเสร็จแล้ว</option>
                </select>
            </div>
            <button
                class="btn btn-square w-16 h-full border border-neutral-100 btn-sm hover:bg-neutral hover:text-neutral-100"
                @click="handleEditTask">แก้ไข</button>
        </div>

    </div>
</template>