<template>
    <ul>
        <li v-for="set in sets" :key="set.id">
            <a
                :href="'/beatmaps/' + set.id"
                :class="id == set.id ? 'bold' : ''"
            >
                {{ set.version }} ({{ set.diff.toFixed(2) }}*)
            </a>
        </li>
    </ul>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["id", "sets"],
});
</script>

<style scoped>
.bold {
    font-weight: bolder;
}
</style>
