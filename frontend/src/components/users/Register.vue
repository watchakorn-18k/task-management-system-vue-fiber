<script setup>
import { ref } from 'vue';
import ErrorUser from '../ErrorUser.vue';
import { registerUsers } from '@/services/userService'
import { useRouter } from 'vue-router'
const router = useRouter();

const userData = ref({
    username: "",
    password: "",
    password_confirm: ""
})
const errorUser = ref({
    errUsername: "",
    errPassword: "",
    errConfirmPassword: ""
});

const handleRegister = async () => {
    if (errorUser.value.errUsername === `${userData.value.username} ถูกใช้ไปแล้ว`) {
        return;
    }
    if (userData.value.username === "") {
        errorUser.value.errUsername = "Username จำเป็นต้องกรอก";
    } else {
        errorUser.value.errUsername = "";
    }
    if (userData.value.password === "") {
        errorUser.value.errPassword = "Password จำเป็นต้องกรอก";
    } else {
        errorUser.value.errPassword = "";
    }
    if (userData.value.password_confirm === "") {
        errorUser.value.errConfirmPassword = "ยืนยัน Password จำเป็นต้องกรอก";
        return;
    }
    if (userData.value.password !== userData.value.password_confirm) {
        errorUser.value.errConfirmPassword = "Password ไม่ตรงกัน";
        console.log(userData.value.password, userData.value.password_confirm);

        return;
    }
    try {
        errorUser.value.errConfirmPassword = "";
        const res = await registerUsers(userData.value);
        if (res === "username already exist") {
            errorUser.value.errUsername = `${userData.value.username} ถูกใช้ไปแล้ว`;
            return;
        }
        router.push("/login");
    } catch (error) {
        console.log(error);
    }
}

const handleUsername = async (e) => {
    errorUser.value.errUsername = "";
}
</script>
<template>
    <div class="flex flex-col items-center w-full mt-16">

        <h1 class="text-2xl font-bold">เข้าสู่ระบบ</h1>
        <div class="flex items-center justify-center w-full">
            <form @submit.prevent="handleRegister"
                class="flex flex-col justify-center gap-2 m-2 w-full sm:w-1/2 md:w-1/4 ">
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-user"></span>
                    <input type="text" class="grow" placeholder="Username" v-model="userData.username"
                        v-on:input="handleUsername" />
                </label>
                <ErrorUser :errorText="errorUser.errUsername" />
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-key"></span>
                    <input type="password" class="grow" placeholder="Password" v-model="userData.password" />
                </label>
                <ErrorUser :errorText="errorUser.errPassword" />
                <label class="input input-bordered flex items-center gap-2">
                    <span class="pi pi-key"></span>
                    <input type="password" class="grow" placeholder="ยืนยัน Password"
                        v-model="userData.password_confirm" />
                </label>
                <ErrorUser :errorText="errorUser.errConfirmPassword" />
                <span class="flex justify-end">
                    <RouterLink to="/login" class="hover:underline hover:text-primary">กลับไปเข้าสู่ระบบ
                    </RouterLink>
                </span>

                <button type="submit" class="btn w-full btn-primary text-white">ลงทะเบียนสมาชิก</button>
            </form>
        </div>

    </div>
</template>