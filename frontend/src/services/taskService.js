
import { config } from '@/config';
import { randomText } from "@/utils/utils";

const urlTasks = config.apiUrl + '/api/task';


const getTasksAll = async () => {
    try {
        const response = await fetch(`${urlTasks}/get_all`);
        const data = await response.json();
        return data.data
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

const editTask = async (taskId, taskData) => {
    try {
        const response = await fetch(`${urlTasks}/edit_task/${taskId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(taskData),
        });
    } catch (err) {
        console.error('Error updating task:', err);
    }
};

export { getTasksAll, addTask, deleteTask, editTask };