
import { config } from '@/config';
import { ref } from 'vue';
import { randomText } from "@/utils";

const urlTasks = config.apiUrl + '/api/task';

const tasks = ref([]);



const getTasksAll = async () => {
    try {
        const response = await fetch(`${urlTasks}/get_all`);
        const data = await response.json();
        tasks.value = data.data;
        console.log(tasks.value);
    } catch (err) {
        console.error('Error fetching tasks:', err);
    }
};

const addTask = async (taskData) => {
    try {
        const taskId = `task-${randomText(4)}`
        const response = await fetch(`${urlTasks}/add_task`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                task_id: taskId,
                name: taskData.name,
                details: taskData.details,
                status: taskData.status ? taskData.status : "ทำอยู่",
            }),
        });
        if (!response.ok) {
            throw new Error('Failed to add task');
        }
        return taskId

    } catch (err) {
        console.error('Error adding new task:', err);
    }
};

const deleteTask = async (taskId) => {
    try {
        const response = await fetch(`${urlTasks}/delete_task/${taskId}`, {
            method: 'DELETE',
        });
    } catch (err) {
        console.error('Error deleting task:', err);
    }
};

export { tasks, getTasksAll, addTask, deleteTask };