<script setup>
import { ref } from 'vue';
import ErrorUser from '../ErrorUser.vue';
import { authLogin } from '@/services/userService'
const userData = ref({
    username: "test_1234",
    password: "test"
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
        const res = await authLogin(userData.value);
        console.log(res);

    } catch (error) {
        console.log(error);

    }
}
</script>
<template>
    <div class="flex flex-col items-center w-full mt-16">
        <h1 class="text-2xl font-bold">เข้าสู่ระบบ</h1>
        <div class="flex items-center justify-center w-full">

            <form @submit.prevent="handleLogin"
                class="flex flex-col justify-center gap-2 m-2 w-full sm:w-1/3 md:w-1/5 ">
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-user"></span>
                    <input type="text" class="grow" placeholder="Username" v-model="userData.username" />
                </label>
                <ErrorUser :errorText="errorUser.errUsername" />
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-key"></span>
                    <input type="password" class="grow" value="password" placeholder="Password"
                        v-model="userData.password" />
                </label>
                <ErrorUser :errorText="errorUser.errPassword" />
                <span class="flex justify-end">
                    <a href="#">สมัครสมาชิก</a>
                </span>
                <button type="submit" class="btn w-full btn-primary text-white">เข้าสู่ระบบ</button>
            </form>
        </div>
    </div>
</template>