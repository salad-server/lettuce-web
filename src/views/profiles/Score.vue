<template>
    <h1>{{ id }}</h1>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["id"],
});
</script>
