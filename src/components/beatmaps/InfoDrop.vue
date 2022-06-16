<template>
    <ul>
        <li v-for="set in sets" :key="set.id">
            <a :href="'/beatmaps/' + set.id"
                >{{ set.version }} ({{ set.diff.toFixed(2) }}*)</a
            >
        </li>
    </ul>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["sets"],
});
</script>
