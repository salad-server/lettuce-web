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

    <div class="buttons">
        <a :href="download" class="button">Download Replay</a>
        <button v-if="myScore" class="button" @click="pin">
            {{ pinned ? "Unpin" : "Pin" }} Score
        </button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { SimpleResponse } from "@/types/general";

import * as m from "@/util/mods";
import * as alert from "@/util/error";

import moment from "moment";
import Swal from "sweetalert2";
import config from "../../../config.json";

export default defineComponent({
    props: ["score"],
    data() {
        return {
            pinned: this.score.pinned,
        };
    },

    computed: {
        prof() {
            const mm = this.score.play_mode.split("!");
            return `/profile/${this.score.uid}?g=${mm[0]}&m=${mm[1]}`;
        },

        pfp() {
            return `${config.api}/users/${this.score.uid}/avatar`;
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

        myScore() {
            return this.score.uid === this.$store.state.id;
        },
    },

    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),

        async pin() {
            const res: SimpleResponse = await fetch(
                `${config.api}/@me/pinned?s=${this.score.id}`,
                {
                    method: this.pinned ? "DELETE" : "POST",
                    headers: {
                        Token: this.$store.state.token,
                    },
                }
            )
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            if (res.code == 200) {
                Swal.fire({
                    title: "Success!",
                    text: `${this.pinned ? "Unpinned" : "Pinned"} score.`,
                    icon: "success",
                });

                this.pinned = !this.pinned;
            } else {
                Swal.fire({
                    title: "Error!",
                    text: "Only passed scores can be pinned/unpinned.",
                    icon: "error",
                });
            }
        },
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
    width: 60%;
    margin: auto;
}

td:nth-child(1) {
    font-weight: bolder;
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

@media screen and (max-width: 727px) {
    table {
        width: 100%;
    }

    .buttons {
        display: none;
    }
}
</style>
