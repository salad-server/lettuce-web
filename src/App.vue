<template>
    <Navbar />
    <router-view />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Navbar from "@/components/global/Navbar.vue";

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
</style>
