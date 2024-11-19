<script setup>
import { stateLogin, removeCookieUserJWT, getCookieUserJWT } from '@/utils/cookies.js';
import { RouterLink, useRoute } from 'vue-router';
import { getProfile } from '@/services/userService';
import { onMounted, ref, onUpdated } from 'vue';

const route = useRoute();

const handleLogout = () => {
    removeCookieUserJWT();
}

const username = ref('');

onMounted(async () => {
    if (stateLogin.isLogin) {
        try {
            const profile = await getProfile();
            username.value = profile.username;

        } catch (error) {
            console.log(error);
        }
    }
})
onUpdated(async () => {
    if (stateLogin.isLogin) {
        try {
            const profile = await getProfile();
            username.value = profile.username;

        } catch (error) {
            console.log(error);
        }
    }
});

</script>

<template>
    <div class="dropdown dropdown-end" v-if="route.path !== '/login'">
        <div class="flex items-center gap-2">
            <p v-if="stateLogin.isLogin">{{ username }}</p>
            <div tabindex="0" role="button"
                class="btn btn-ghost btn-circle avatar  border border-neutral hover:bg-neutral">
                <div class="p-1 rounded-full">
                    <span class="pi pi-user text-xl"></span>
                </div>
            </div>

        </div>
        <ul tabindex="0"
            class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] border border-neutral mt-1 w-52 p-2 shadow shadow-base-200">
            <li v-if="stateLogin.isLogin">
                <a class="justify-between hover:bg-base-100 focus:bg-base-100">
                    โปรไฟล์
                    <span class="badge text-xs">เร็วๆนี้</span>
                </a>
            </li>
            <li v-if="!stateLogin.isLogin">
                <RouterLink to="/login" class="justify-between">
                    เข้าสู่ระบบ
                </RouterLink>
            </li>
            <li v-if="stateLogin.isLogin"><a @click="handleLogout">ออกจากระบบ</a></li>
        </ul>
    </div>
</template>