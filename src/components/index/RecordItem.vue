<template>
    <table v-if="scoreExists">
        <tr>
            <td>
                <img :src="pfp" width="50" height="50" />
            </td>
            <td>
                <h1>{{ score.rank }}</h1>
            </td>
            <td>
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
                        <router-link :to="pflink">{{
                            score.username
                        }}</router-link>
                    </div>

                    <div>
                        <b>{{ score.play_mode }}</b>
                    </div>
                </div>
            </td>
        </tr>
    </table>
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
