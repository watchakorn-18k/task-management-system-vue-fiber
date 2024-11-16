<script setup>
defineProps({
    errorText: {
        type: String
    }
})

</script>
<template>
    <p class="text-sm text-red-500 mx-2" v-if="errorText">{{ errorText }}</p>
</template>