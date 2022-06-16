<template>
    <div v-if="data">
        <table>
            <ScoreItem
                v-for="score in data"
                :key="score.id"
                :score="score"
                :mode="mode"
            />
        </table>

        <button @click="this.$emit('clicked')">Load</button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ScoreItem from "./ScoreItem.vue";

export default defineComponent({
    props: ["data", "mode"],
    components: { ScoreItem },
});
</script>
