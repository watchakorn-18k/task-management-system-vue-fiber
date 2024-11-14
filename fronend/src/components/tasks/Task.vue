<script setup>
import Card from './Card.vue';
import { ref } from 'vue';

const tasks = ref([{
    taskId: 1,
    name: "Lorem ipsum1",
    details: "Lorem ipsum dolor sit amet consectetur adipiscing elit",
    status: "ทำอยู่"
},
{
    taskId: 2,
    name: "Lorem ipsum2",
    details: "Lorem ipsum dolor sit amet consectetur adipiscing elit",
    status: "ทำเสร็จแล้ว"
}

]);

</script>

<template>
    <div class="flex flex-wrap gap-4 justify-center">
        <span v-for="task in tasks" :key="task.taskId">
            <Card :title="task.name" :details="task.details" v-model:status="task.status" />
        </span>
    </div>


</template>