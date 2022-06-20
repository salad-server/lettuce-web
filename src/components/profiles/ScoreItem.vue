<template>
    <tr :style="bg">
        <td class="td-rank">
            <h1
                class="rank is-size-2"
                :class="'rank-' + score.rank.toLowerCase()"
            >
                {{ score.rank }}
            </h1>
        </td>
        <td class="td-info">
            <router-link :to="'/beatmaps/' + score.map.map_id">
                <b>{{ score.map.title }} [{{ score.map.version }}]</b>
            </router-link>
            <div class="td-m">
                <b>{{ comma(score.score) }} | </b>
                <span>
                    {{ comma(score.c300) }} / {{ comma(score.c100) }} /
                    {{ comma(score.c50) }} / {{ comma(score.miss) }} |
                    {{ comma(score.combo) }}x
                </span>

                <span v-if="showMap">
                    <span> | </span>
                    <span :class="status">{{ score.map.status }}</span>
                    <span v-if="showStatus"> / Unsubmitted </span>
                </span>
            </div>
            <div class="td-m-s">
                + {{ modsTxt }} ({{ score.acc.toFixed(2) }}%)
            </div>
            <div :title="score.play_date">
                {{ time(score.play_date) }}
            </div>
        </td>
        <td class="td-mods">
            <div v-html="mods"></div>
        </td>
        <td class="td-stats">
            <h3>{{ score.pp.toFixed(2) }}pp</h3>
            <router-link
                v-if="score.submission_status"
                :to="`/score/${score.id}`"
                >More Info</router-link
            >
        </td>
    </tr>
    <tr class="spacer"></tr>
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

        modsTxt() {
            return m
                .modstr(
                    m.convertMods(
                        this.score.mods,
                        ["std", "taiko", "catch", "mania"].indexOf(this.mode)
                    )
                )
                .join(", ");
        },

        showStatus() {
            return this.score.submission_status != 2;
        },

        showMap() {
            return this.score.map.status != "Ranked";
        },

        status() {
            return `status-${this.score.map.status.toLowerCase()}`;
        },

        bg() {
            return `
                background-image: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.6)), url("https://assets.ppy.sh/beatmaps/${this.score.map.set_id}/covers/cover.jpg");
            `;
        },
    },
});
</script>

<style scoped>
.spacer {
    height: 3px;
}

tr {
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
}

td {
    padding: 9px;
}

.td-rank *,
.td-stats * {
    display: block;
    text-align: center;
    margin: auto;
}

.td-stats h3 {
    font-weight: bolder;
    color: #fff;
}

.td-stats a {
    line-height: 40px;
    font-size: 0.8em;
}

.td-mods * {
    text-align: right;
}

.td-info {
    color: rgba(255, 255, 255, 0.7);
}

.td-info a {
    color: #fff;
}

.td-info a:hover,
.td-stats a:hover {
    color: rgba(255, 255, 255, 0.7);
    text-decoration: underline;
}

.td-rank {
    width: 3%;
}

.td-mods {
    width: 10%;
    vertical-align: middle;
}

.td-stats {
    width: 10%;
}

.td-m-s {
    display: none;
    color: #fff;
    font-weight: bolder;
}

@media screen and (max-width: 768px) {
    .td-m,
    .td-mods {
        display: none !important;
    }

    .td-m-s {
        display: block !important;
    }
}
</style>
