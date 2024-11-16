<script setup>
import { defineEmits } from 'vue';
const emit = defineEmits(['filterStatusEmit']);

const handleFilter = (event) => {
    emit('filterStatusEmit', event);
}
</script>

<template>
    <select class="select select-bordered join-item" @change="handleFilter">
        <option>ทั้งหมด</option>
        <option>ทำอยู่</option>
        <option>ทำเสร็จแล้ว</option>
    </select>
</template>