<template>
    <tr v-if="score">
        <td>#{{ rank }}</td>
        <td>{{ comma(score.score) }}</td>
        <td>{{ score.acc.toFixed(2) }}%</td>
        <td>
            <router-link :to="prof">
                <span>{{ score.user.country }}</span> |
                <span>{{ score.user.username }}</span>
            </router-link>
        </td>
        <td>{{ comma(score.max_combo) }}</td>
        <td>{{ comma(score.c300) }}</td>
        <td>{{ comma(score.c100) }}</td>
        <td>{{ comma(score.c50) }}</td>
        <td>{{ comma(score.miss) }}</td>
        <td>{{ score.pp.toFixed(2) }}pp</td>
        <td>{{ time }}</td>
        <td v-html="mods"></td>
        <td>
            <router-link :to="'/score/' + score.id">Info</router-link>
        </td>
    </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";
import * as m from "@/util/mods";

export default defineComponent({
    props: ["score", "rank", "mod", "mode"],
    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },

    computed: {
        time() {
            return moment(this.score.play_date).fromNow();
        },

        mods() {
            const mode = ["std", "taiko", "catch", "mania"].indexOf(this.mode);
            const mods = m.convertMods(this.score.mods, mode);
            const str = m.modstr(mods);

            // prettier-ignore
            return str.map((s) => `<img src="/img/mods/${s.toLowerCase()}.png" alt="${s}" />`).join("") || "None";
        },

        prof() {
            return `/profile/${this.score.user.id}?g=${this.mod}&m=${this.mode}`;
        },
    },
});
</script>
