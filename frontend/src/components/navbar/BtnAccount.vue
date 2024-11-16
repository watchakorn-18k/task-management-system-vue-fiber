<script setup>

</script>

<template>
    <div class="dropdown dropdown-end">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar  border border-neutral hover:bg-neutral">
            <div class="p-1 rounded-full">
                <span class="pi pi-user text-xl"></span>
            </div>
        </div>
        <ul tabindex="0"
            class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] border border-neutral mt-1 w-52 p-2 shadow shadow-base-200">
            <li>
                <a class="justify-between">
                    Profile
                    <span class="badge">New</span>
                </a>
            </li>
            <li><a>Settings</a></li>
            <li><a>Logout</a></li>
        </ul>
    </div>
</template>