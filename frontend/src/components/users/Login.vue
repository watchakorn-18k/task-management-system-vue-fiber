<script setup>
import { ref } from 'vue';
import ErrorUser from '../ErrorUser.vue';
import { authLogin } from '@/services/userService'
import { setCookieUserJWT } from '@/utils/cookies.js';
import { useRouter } from 'vue-router'
import { state } from '@/utils/utils.js';

const router = useRouter();


const userData = ref({
    username: "",
    password: ""
})
const errorUser = ref({
    errUsername: "",
    errPassword: ""
});
const handleLogin = async () => {
    if (userData.value.username === "") {
        errorUser.value.errUsername = "Username จำเป็นต้องกรอก";
    } else {
        errorUser.value.errUsername = "";
    }
    if (userData.value.password === "") {
        errorUser.value.errPassword = "Password จำเป็นต้องกรอก";
        return;
    }
    try {
        state.isLoading = true;
        const res = await authLogin(userData.value);
        if (res === "password not match") {
            errorUser.value.errPassword = "Password ไม่ถูกต้อง";
            state.isLoading = false;
            return;
        }
        if (res === "username not found") {
            errorUser.value.errUsername = "ไม่พบ Username ในระบบ";
            state.isLoading = false;
            return;
        }
        setCookieUserJWT(res);
        state.isLoading = false;
        router.push("/");

    } catch (error) {
        state.isLoading = false;
        console.log(error);
    }
}
</script>
<template>
    <div class="flex flex-col items-center w-full mt-16">

        <h1 class="text-2xl font-bold">เข้าสู่ระบบ</h1>
        <div class="flex items-center justify-center w-full">
            <form @submit.prevent="handleLogin"
                class="flex flex-col justify-center gap-2 m-2 w-full sm:w-1/2 md:w-1/4 ">
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-user"></span>
                    <input type="text" class="grow" placeholder="Username" v-model="userData.username"
                        :disabled="state.isLoading" />
                </label>
                <ErrorUser :errorText="errorUser.errUsername" />
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-key"></span>
                    <input type="password" class="grow" placeholder="Password" v-model="userData.password"
                        :disabled="state.isLoading" />
                </label>
                <ErrorUser :errorText="errorUser.errPassword" />
                <span class="flex justify-end">
                    <RouterLink to="/register" class="hover:underline hover:text-primary">สมัครสมาชิก
                    </RouterLink>
                </span>
                <button type="submit" class="btn w-full btn-primary text-white" :disabled="state.isLoading">เข้าสู่ระบบ
                    <span class="loading loading-spinner loading-md" v-if="state.isLoading"></span></button>
            </form>
        </div>

    </div>
</template>