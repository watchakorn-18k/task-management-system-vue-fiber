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
const emit = defineEmits(['update:status', 'remove-item']);
const isOpen = ref(false);

const updateStatus = (newStatus) => {
    isOpen.value = false;
    emit('update:status', newStatus); // ใช้งาน emit ที่ถูกต้อง
};

const syncOpenState = (e) => {
    isOpen.value = e.target.open;
};

const deleteCard = () => {
    emit('remove-item');
}

</script>
<template>
    <div class="card bg-neutral-100 w-80 shadow-xl">
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

        <div class="flex justify-end items-center">
            <div class="flex flex-row items-center">
                <p>สถานะ</p>

                <details className="dropdown" :open="isOpen" @toggle="syncOpenState" autofocus>
                    <summary :class="{
                        'hover:bg-primary': status === 'ทำอยู่',
                        'hover:bg-success': status === 'ทำเสร็จแล้ว',
                        'badge-primary': status === 'ทำอยู่',
                        'badge-success': status === 'ทำเสร็จแล้ว',

                    }" class="badge badge-outline m-4 hover:text-neutral-100 p-4 cursor-pointer">{{ status }}</summary>
                    <ul
                        className="menu dropdown-content bg-neutral-100 rounded-box z-[1] w-36 p-2 shadow border-neutral border-[1px]">
                        <li @click="updateStatus('ทำอยู่')"><a class="text-primary">ทำอยู่</a></li>
                        <li><a class="text-success" @click="updateStatus('ทำเสร็จแล้ว')">ทำเสร็จแล้ว</a></li>
                    </ul>
                </details>
            </div>
        </div>

    </div>
</template>