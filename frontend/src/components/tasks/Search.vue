<script setup>

const emit = defineEmits(['searchTaskEmit']);

const handleSearchTask = (event) => {
    emit('searchTaskEmit', event);
}
</script>

<template>
    <label class="input input-bordered flex w-full md:w-full items-center gap-2">
        <input type="text" class="grow" placeholder="ค้นหา" @input="handleSearchTask" />
        <span class="pi pi-search"></span>
    </label>
</template>