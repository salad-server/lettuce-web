<template>
    <Navbar />
    <router-view />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Navbar from "@/components/global/Navbar.vue";

import "@/assets/main.scss";

export default defineComponent({
    name: "HomeView",
    components: {
        Navbar,
    },
});
</script>

<style>
.enabled {
    color: red;
}

.td-rank {
    width: 0.1%;
    white-space: nowrap;
}
</style>
