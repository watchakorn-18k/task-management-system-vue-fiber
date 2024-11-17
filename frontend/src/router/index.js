import { createRouter, createWebHistory } from 'vue-router'
import TasksView from '@/views/TasksView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import { stateLogin } from '@/utils/cookies.js';
const routes = [
    { path: '/', component: TasksView },
    { path: '/login', component: LoginView },
    { path: '/register', component: RegisterView },
    { path: '/:catchAll(.*)', redirect: '/' },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const guestOnlyRoutes = ['/login', '/register'];
    if (guestOnlyRoutes.includes(to.path) && stateLogin.isLogin) {
        next('/'); // Redirect ไปที่หน้า Home
    } else {
        next(); // ให้ผ่านไป
    }
});

export default router
