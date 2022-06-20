<template>
    <div class="column is-9 has-text-centered">
        <div>
            <!-- prettier-ignore -->
            <div class="my-1">
                <button :class="clicked('vn')" v-show="!enabled('vn')" @click="updateMod('vn')">vn</button>
                <button :class="clicked('rx')" v-show="!enabled('rx')" @click="updateMod('rx')">rx</button>
                <button :class="clicked('ap')" v-show="!enabled('ap')" @click="updateMod('ap')">ap</button>
            </div>

            <!-- prettier-ignore -->
            <div class="my-1" v-if="canSwitch">
                <button :class="clicked('std')" v-show="!enabled('std')" @click="updateMode('std')">std</button>
                <button :class="clicked('taiko')" v-show="!enabled('taiko')" @click="updateMode('taiko')">taiko</button>
                <button :class="clicked('catch')" v-show="!enabled('catch')" @click="updateMode('catch')">catch</button>
                <button :class="clicked('mania')" v-show="!enabled('mania')" @click="updateMode('mania')">mania</button>
            </div>
        </div>

        <div v-if="anyscores" class="table-container">
            <table class="table">
                <tr>
                    <td>Rank</td>
                    <td>Score</td>
                    <td>Accuracy</td>
                    <td>Player</td>
                    <td>Max Combo</td>
                    <td>300</td>
                    <td>100</td>
                    <td>50</td>
                    <td>Miss</td>
                    <td>PP</td>
                    <td>Date</td>
                    <td>Mods</td>
                    <td>Info</td>
                </tr>

                <LeaderboardItem
                    v-for="(score, i) in scores"
                    :key="score.id"
                    :score="score"
                    :rank="i + 1"
                    :mod="mod"
                    :mode="mode"
                />
            </table>
            <button class="button" @click="loadScores()">Load more</button>
        </div>
        <h2 v-else class="subtitle is-size-4 my-5">No scores yet :(</h2>
    </div>
</template>

<script lang="ts">
import { BeatmapScore } from "@/types/beatmap";
import { defineComponent } from "vue";
import * as alert from "@/util/error";

import config from "../../../config.json";
import LeaderboardItem from "./LeaderboardItem.vue";

export default defineComponent({
    props: ["id", "play_mode"],
    components: { LeaderboardItem },

    data() {
        return {
            canSwitch: this.play_mode == "vn!std",
            page: 0,
            scores: [] as BeatmapScore[],
            mod: "",
            mode: "",
            disabled: {
                mod: [""],
                mode: [""],
            },
        };
    },

    async mounted() {
        this.mod = this.play_mode.split("!")[0];
        this.mode = this.play_mode.split("!")[1];

        await this.loadScores();
    },

    methods: {
        async loadScores() {
            const res: BeatmapScore[] = await fetch(
                `${config.api}/beatmap/${this.id}/leaderboard?m=${this.mod}!${
                    this.mode
                }&p=${this.page++}`
            )
                .then((j) => j.json())
                .catch(() => alert.API());

            if (!res) return;
            const ids = this.scores.map((s: BeatmapScore) => s.id);

            res.forEach((r: BeatmapScore) => {
                if (!ids.includes(r.id)) {
                    this.scores.push(r);
                }
            });
        },

        all() {
            this.page = 0;
            this.scores.length = 0;
            this.loadScores();
        },

        updateMod(m: string) { this.mod = m; }, // prettier-ignore
        updateMode(m: string) { this.mode = m; }, // prettier-ignore
        enabled(m: string) {
            return [...this.disabled.mode, ...this.disabled.mod].includes(m);
        },

        clicked(m: string) {
            return [this.mod, this.mode].includes(m)
                ? "is-primary button"
                : "button";
        },
    },

    watch: {
        mod(n: string) {
            switch (n) {
                case "vn":
                    this.disabled.mod = [];
                    break;
                case "rx":
                    this.disabled.mod = ["mania"];
                    break;
                case "ap":
                    this.disabled.mod = ["taiko", "catch", "mania"];
                    break;
            }

            this.all();
        },

        mode(n: string) {
            switch (n) {
                case "std":
                    this.disabled.mode = [];
                    break;
                case "taiko":
                case "catch":
                    this.disabled.mode = ["ap"];
                    break;
                case "mania":
                    this.disabled.mode = ["rx", "ap"];
                    break;
            }

            this.all();
        },
    },

    computed: {
        anyscores() {
            return this.scores.length > 0;
        },
    },
});
</script>

<style scoped>
table {
    width: 100%;
    text-align: left;
}

span {
    line-height: 40px;
}
</style>
