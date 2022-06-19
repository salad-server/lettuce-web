<template>
    <tr>
        <td>
            <h1>{{ score.rank }}</h1>
        </td>
        <td>
            <router-link :to="'/beatmaps/' + score.map.map_id">
                <b>{{ score.map.title }} [{{ score.map.version }}]</b>
            </router-link>
            <div>
                <b>{{ comma(score.score) }} | </b>
                <span>
                    {{ comma(score.c300) }} / {{ comma(score.c100) }} /
                    {{ comma(score.c50) }} / {{ comma(score.miss) }} |
                    {{ comma(score.combo) }}x
                </span>
            </div>
            <div>
                <span v-if="showMap">{{ score.map.status }}</span>
                <span v-if="showMap && showStatus"> / </span>
                <span v-if="showStatus">Unsubmitted</span>
            </div>
            <div :title="score.play_date">
                {{ time(score.play_date) }}
            </div>
        </td>
        <td>
            <div v-html="mods"></div>
        </td>
        <td>
            <h3>{{ score.pp.toFixed(2) }}pp</h3>
            <router-link
                v-if="score.submission_status"
                :to="`/score/${score.id}`"
                >More Info</router-link
            >
        </td>
    </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";
import * as m from "@/util/mods";

export default defineComponent({
    props: ["score", "mode"],
    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),

        time(date: Date) {
            return moment(new Date(date)).fromNow();
        },
    },

    computed: {
        mods() {
            const mode = ["std", "taiko", "catch", "mania"].indexOf(this.mode);
            const mods = m.convertMods(this.score.mods, mode);
            const str = m.modstr(mods);

            // prettier-ignore
            return str.map((s) => `<img src="/img/mods/${s.toLowerCase()}.png" alt="${s}" />`).join("");
        },

        showStatus() {
            return this.score.submission_status != 2;
        },

        showMap() {
            return this.score.map.status != "Ranked";
        },
    },
});
</script>
