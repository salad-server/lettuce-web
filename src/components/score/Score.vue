<template>
    <table>
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
            <td>{{ score.rank }}</td>
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
            return str.map((s) => `<img src="/img/${s.toLowerCase()}.png" alt="${s}" />`).join("") || "None";
        },

        date() {
            const t = new Date(this.score.play_date);
            const d = moment(t);

            return `${d.format("DD/MM/yy - hh:mm.ssa")} (${d.fromNow()})`;
        },
    },

    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },
});
</script>
