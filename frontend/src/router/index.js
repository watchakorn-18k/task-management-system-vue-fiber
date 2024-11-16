import { createRouter, createWebHistory } from 'vue-router'
import TasksView from '@/views/TasksView.vue'
import LoginView from '@/views/LoginView.vue'

const routes = [
    { path: '/', component: TasksView },
    { path: '/login', component: LoginView },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
