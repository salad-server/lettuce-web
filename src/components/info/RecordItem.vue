<template>
    <div v-if="scoreExists" class="grid-item columns">
        <div class="column is-2">
            <img :src="pfp" width="70" height="70" />
        </div>
        <div class="column is-1">
            <h1 class="rank">{{ score.rank }}</h1>
        </div>
        <div class="column">
            <router-link :to="'/beatmaps/' + score.map.map_id">
                {{ score.map.title }} [{{ score.map.version }}]
            </router-link>

            <div>
                <span>
                    {{ score.map.diff.toFixed(2) }}* |
                    {{ score.pp.toFixed(2) }}pp |
                    {{ comma(score.score) }}
                </span>

                <div style="float: right">
                    <router-link :to="pflink">{{ score.username }}</router-link>
                </div>

                <div>
                    <b>{{ score.play_mode }}</b>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["score"],
    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },

    computed: {
        pfp() {
            return `${config.avatar}/${this.score.uid}`;
        },

        scoreExists() {
            return this.score.id != 0;
        },

        pflink() {
            const g = this.score.play_mode.split("!");

            return `/profile/${this.score.uid}?g=${g[0]}&m=${g[1]}`;
        },
    },
});
</script>

<style scoped>
.column img {
    display: block;
    margin: auto;
}

.is-1 {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 30px;
    font-weight: bolder;
    max-height: 100px;
}
</style>
