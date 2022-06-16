<template>
    <div v-if="data && scoresExist">
        <table>
            <ScoreItem
                v-for="score in data"
                :key="score.id"
                :score="score"
                :mode="mode"
            />
        </table>

        <button @click="$emit('clicked')">Load</button>
    </div>
    <h3 v-else>No scores yet!</h3>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ScoreItem from "./ScoreItem.vue";

export default defineComponent({
    props: ["data", "mode"],
    components: { ScoreItem },
    computed: {
        scoresExist() {
            return !!this.data.length;
        },
    },
});
</script>
