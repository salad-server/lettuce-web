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

        <button class="button" @click="$emit('clicked')">Load More</button>
    </div>
    <h3 v-else class="has-text-centered is-size-4 title addtop">
        No scores yet!
    </h3>
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

<style scoped>
table {
    width: 100%;
}

.button {
    display: block;
    margin: 0 auto;
}

.addtop {
    margin: 50px;
    color: rgba(0, 0, 0, 0.7);
}

@media screen and (max-width: 768px) {
    button {
        width: 80%;
    }
}
</style>
