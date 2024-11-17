<script setup>
import Navbar from './components/navbar/Navbar.vue';
import { RouterView } from 'vue-router';
import { stateLogin, getCookieUserJWT } from '@/utils/cookies.js';
import { onMounted } from 'vue';

onMounted(async () => {
    const jwt = getCookieUserJWT();
    if (jwt) { 
        stateLogin.isLogin = true;
        stateLogin.jwt = getCookieUserJWT();
    }
})
</script>
<template>
  <Navbar />
  <RouterView />
</template>