<template>
    <Error v-if="error" :code="errCode" :msg="errorMsg" />
    <div v-else class="columns is-centered">
        <div class="column dcard is-8 has-text-centered">
            <div class="my-5">
                <div>
                    <button
                        v-for="m in playmodes"
                        :class="isMode(m)"
                        :key="m"
                        @click="mode = m"
                    >
                        {{ m }}
                    </button>
                </div>

                <h1 class="is-size-1 my-5">Player Leaderboard</h1>
                <div>
                    <button :class="isSort(0)" @click="sort = 0">PP</button>
                    <button :class="isSort(1)" @click="sort = 1">Score</button>
                </div>
            </div>

            <LoadingMini v-if="loading" />
            <div v-else>
                <table v-if="anyplayers" class="table">
                    <tr>
                        <td>Rank</td>
                        <td>User</td>
                        <td>PP</td>
                        <td>Total Score</td>
                        <td>Ranked Score</td>
                        <td>Playcount</td>
                    </tr>

                    <Rank
                        v-for="(rank, i) in users"
                        :key="rank.user.id"
                        :rank="rank"
                        :page="page + i"
                        :mode="mode"
                    />
                </table>
                <div v-else>
                    <img
                        src="/img/errors/no-players.png"
                        alt="No players"
                        width="400"
                    />
                    <h4 class="is-size-3">No players here...</h4>
                </div>
            </div>

            <div class="my-5">
                <button class="button" @click="updatePage(-1)">&lt;</button>
                <span class="is-size-6 mx-3">{{ page + 1 }}</span>
                <button class="button" @click="updatePage(1)">&gt;</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Leaderboard } from "@/types/ranking";

import * as alert from "@/util/error";
import router from "@/router";

import Error from "@/components/global/Error.vue";
import LoadingMini from "@/components/global/LoadingMini.vue";
import Rank from "@/components/leaderboard/Rank.vue";
import config from "../../../config.json";

const modes = [
    "vn!std",
    "vn!taiko",
    "vn!catch",
    "vn!mania",
    "rx!std",
    "rx!taiko",
    "rx!catch",
    "ap!std",
];

export default defineComponent({
    components: { LoadingMini, Error, Rank },
    data() {
        return {
            mode: "",
            sort: 0,
            page: 0,
            users: [] as Leaderboard[],

            loading: true,
            playmodes: modes,

            error: false,
            errorMsg: "",
            errCode: 0,
        };
    },

    async mounted() {
        const mode = new URLSearchParams(window.location.search).get("m") || "";

        if (mode != "" && !modes.includes(mode)) {
            this.loading = false;
            this.error = true;
            this.errorMsg = "Invalid gamemode!";
            this.errCode = 400;

            return;
        }

        this.mode = mode || "vn!std";
        await this.loadAll(true);
    },

    methods: {
        async loadLB() {
            const res: Leaderboard[] = await fetch(
                `${config.api}/leaderboard?m=${this.mode}&p=${this.page}&s=${this.sort}`
            )
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            const ids = this.users.map((u: Leaderboard) => u.user.id);

            for (const lb of res) {
                if (!ids.includes(lb.user.id)) {
                    this.users.push(lb);
                }
            }

            this.loading = false;
        },

        async loadAll(p: boolean) {
            if (p) {
                this.page = 0;
            }

            this.users.length = 0;
            this.loading = true;

            await this.loadLB();
        },

        isMode(what: string) {
            return this.mode == what ? "is-primary button" : "button";
        },

        isSort(what: number) {
            return this.sort == what ? "is-primary button" : "button";
        },

        updatePage(n: number) {
            if (this.page + n <= -1) {
                return;
            }

            this.page += n;
            this.loadAll(false);
        },
    },

    computed: {
        anyplayers(): boolean {
            return this.users.length > 0;
        },
    },

    watch: {
        mode() {
            this.loadAll(true);
            router.push({ query: { m: this.mode } });
        },

        sort() {
            this.loadAll(true);
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
