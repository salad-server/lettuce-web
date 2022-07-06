<template>
    <h2>
        <router-link class="title" :to="'/beatmaps/' + score.map.map_id">
            {{ score.map.artist }} - {{ score.map.title }} [{{
                score.map.version
            }}]
        </router-link>
    </h2>
    <table class="table">
        <tr>
            <td>Gamemode</td>
            <td>{{ score.play_mode }}</td>
        </tr>
        <tr>
            <td>PP/Score</td>
            <td>{{ comma(score.pp) }}pp / {{ comma(score.score) }}</td>
        </tr>
        <tr>
            <td>Accuracy</td>
            <td>{{ score.acc.toFixed(2) }}%</td>
        </tr>
        <tr>
            <td>Combo</td>
            <td>
                {{ score.combo }}/{{ score.map.max_combo }}
                <span v-if="score.perfect_score">perfect!</span>
            </td>
        </tr>
        <tr>
            <td>Mods</td>
            <td v-html="mods"></td>
        </tr>
        <tr>
            <td>300 Count</td>
            <td>{{ score.c300 }}x</td>
        </tr>
        <tr>
            <td>100 Count</td>
            <td>{{ score.c100 }}x</td>
        </tr>
        <tr>
            <td>50 Count</td>
            <td>{{ score.c50 }}x</td>
        </tr>
        <tr>
            <td>Geki Count</td>
            <td>{{ score.geki }}x</td>
        </tr>
        <tr>
            <td>Katu Count</td>
            <td>{{ score.katu }}x</td>
        </tr>
        <tr>
            <td>Miss Count</td>
            <td>{{ score.miss }}x</td>
        </tr>
        <tr>
            <td>Grade</td>
            <td :class="'rank rank-' + score.rank.toLowerCase()">
                {{ score.rank }}
            </td>
        </tr>
        <tr>
            <td>Date:</td>
            <td :title="score.play_date">{{ date }}</td>
        </tr>
    </table>

    <b>
        Played by:
        <!-- prettier-ignore -->
        <router-link :to="prof">
            <img :src="pfp" alt="Profile Picture" width="30" height="30" /> {{ score.username }}
        </router-link>
    </b>

    <a :href="download" class="button">Download Replay</a>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import * as m from "@/util/mods";
import moment from "moment";
import config from "../../../config.json";

export default defineComponent({
    props: ["score"],
    computed: {
        prof() {
            const mm = this.score.play_mode.split("!");
            return `/profile/${this.score.uid}?g=${mm[0]}&m=${mm[1]}`;
        },

        pfp() {
            return `${config.avatar}/${this.score.uid}`;
        },

        mods() {
            // prettier-ignore
            const mode = ["std", "taiko", "catch", "mania"].indexOf(this.score.play_mode.substr(3));
            const mods = m.convertMods(this.score.mods, mode);
            const str = m.modstr(mods);

            // prettier-ignore
            return str.map((s) => `<img src="/img/mods/${s.toLowerCase()}.png" alt="${s}" />`).join("") || "None";
        },

        date() {
            const t = new Date(this.score.play_date);
            const d = moment(t);

            return `${d.format("DD/MM/yy - hh:mm.ssa")} (${d.fromNow()})`;
        },

        download() {
            return `${config.cho}/get_replay?id=${this.score.id}&include_headers=true`;
        },
    },

    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },
});
</script>

<style scoped>
b {
    float: right;
}

b img {
    display: inline;
    vertical-align: middle;
    margin: 0 5px;
    width: 1.3em;
    height: 1.3em;
}

.title {
    display: block;
    margin-bottom: 30px;
    text-align: center;
    color: #fff;
}

.title:hover {
    color: rgba(255, 255, 255, 0.7);
    text-decoration: underline;
}

table {
    background-color: transparent;
    color: #fff;
    width: 80%;
    margin: auto;
}

.button {
    background: transparent;
    color: #fff;
    font-weight: bolder;
}

.button:hover {
    color: #000 !important;
    background: #fff;
}

.button:active {
    position: relative;
    top: 3px;
}
</style>
