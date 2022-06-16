<template>
    <div v-if="data">
        <table v-if="scoresExist">
            <ScoreItem
                v-for="score in data"
                :key="score.id"
                :score="score"
                :mode="mode"
            />
        </table>
        <h3 v-else>No scores yet!</h3>

        <button @click="$emit('clicked')">Load</button>
    </div>
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
