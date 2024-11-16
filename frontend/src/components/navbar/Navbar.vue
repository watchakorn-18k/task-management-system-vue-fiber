<script setup>
import logo from '@/assets/img/logo.png';
import BtnAccount from './BtnAccount.vue';
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="navbar bg-base-300">
        <div class="flex-1">
            <RouterLink class="btn btn-ghost" href="/"><img :src="logo" class="w-6 h-6">Task Management System
            </RouterLink>
        </div>
        <BtnAccount />
    </div>
</template>