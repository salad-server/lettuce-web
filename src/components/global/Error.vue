<template>
    <div>
        <h1>Request Error!</h1>
        <h2>{{ msg }}</h2>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["msg"],
});
</script>
