<template>
    <Loading v-if="loading" />
    <div v-else-if="!error" class="columns is-mobile is-centered">
        <div class="column is-11">
            <div class="columns">
                <div class="column is-3">
                    <Info :id="id" :info="info" />

                    <div class="has-text-centered my-3">
                        <!-- prettier-ignore -->
                        <div class="my-1">
                            <button :class="clicked('vn')" :disabled="enabled('vn')" @click="updateMod('vn')">vn</button>
                            <button :class="clicked('rx')" :disabled="enabled('rx')" @click="updateMod('rx')">rx</button>
                            <button :class="clicked('ap')" :disabled="enabled('ap')" @click="updateMod('ap')">ap</button>
                        </div>

                        <!-- prettier-ignore -->
                        <div class="my-1">
                            <button :class="clicked('std')" :disabled="enabled('std')" @click="updateMode('std')">std</button>
                            <button :class="clicked('taiko')" :disabled="enabled('taiko')" @click="updateMode('taiko')">taiko</button>
                            <button :class="clicked('catch')" :disabled="enabled('catch')" @click="updateMode('catch')">catch</button>
                            <button :class="clicked('mania')" :disabled="enabled('mania')" @click="updateMode('mania')">mania</button>
                        </div>
                    </div>

                    <Stats :stats="stats" />
                </div>

                <div class="column">
                    <div v-if="info.bio">
                        <div class="divider">About</div>
                        <p>{{ info.bio }}</p>
                    </div>

                    <div v-if="pinned.length">
                        <div class="divider">Pinned Scores</div>
                        <Scores
                            :data="pinned"
                            @clicked="loadPinned()"
                            :mode="mode"
                        />
                    </div>

                    <div>
                        <div class="divider">Best Scores</div>
                        <Scores
                            :data="best"
                            @clicked="loadBest()"
                            :mode="mode"
                        />
                    </div>

                    <div>
                        <div class="divider">Recent Scores</div>
                        <Scores
                            :data="recent"
                            @clicked="loadRecent()"
                            :mode="mode"
                        />
                    </div>
                </div>
            </div>
        </div>
    </div>
    <Error v-else :msg="errorMsg" />
</template>

<script lang="ts">
import router from "@/router";
import { defineComponent } from "vue";
import { Score } from "@/types/scores";
import { Stats as UserStats, Info as UserInfo } from "@/types/user";

import Loading from "@/components/global/Loading.vue";
import Error from "@/components/global/Error.vue";
import Info from "@/components/profiles/Info.vue";
import Scores from "@/components/profiles/Scores.vue";
import Stats from "@/components/profiles/Stats.vue";

import * as alert from "@/util/error";
import config from "../../../config.json";

type Mode = "std" | "taiko" | "catch" | "mania";
type Mod = "vn" | "rx" | "ap";

function invalidModes(mode: Mode, mod: Mod) {
    const allowed = {
        std: ["vn", "rx", "ap"],
        taiko: ["vn", "rx"],
        catch: ["vn", "rx"],
        mania: ["vn"],
    };

    const properMode = !["std", "taiko", "catch", "mania"].includes(mode);
    const properMod = !["vn", "rx", "ap"].includes(mod);

    return properMode || properMod || !allowed[mode].includes(mod);
}

export default defineComponent({
    props: ["id"],
    components: { Loading, Error, Info, Scores, Stats },
    data() {
        return {
            mod: "",
            mode: "",
            loading: true,

            best: [] as Score[],
            recent: [] as Score[],
            pinned: [] as Score[],
            stats: {} as UserStats,
            info: {} as UserInfo,
            page: {
                best: 0,
                recent: 0,
                pinned: 0,
            },

            disabled: {
                mode: [""],
                mod: [""],
            },

            error: false,
            errorMsg: "",
        };
    },

    async mounted() {
        const mode = new URLSearchParams(window.location.search).get("m") || "";
        const mod = new URLSearchParams(window.location.search).get("g") || "";

        if (mode != "" && mod != "" && invalidModes(mode as Mode, mod as Mod)) {
            this.loading = false;
            this.error = true;
            this.errorMsg = "Invalid mode/mod combo!";

            return;
        }

        this.mode = mode || "std";
        this.mod = mod || "vn";

        await this.loadInfo();
        this.updateEverything();
        this.loading = false;
    },

    methods: {
        async loadInfo() {
            // prettier-ignore
            const res: UserInfo = await fetch(`${config.api}/users/${this.id}`).then((j) => j.json()).catch(() => {
                this.error = true;
                this.errorMsg = "Error loading profile!";
            });

            if (!res) return;
            if (res.id == 0 || [400, 500].includes(res?.code || 0)) {
                this.error = true;
                this.errorMsg = "Profile not found!";
                return;
            }

            this.info = res;
        },

        async loadStats() {
            // prettier-ignore
            const res: UserStats = await fetch(`${config.api}/users/${this.id}/stats?m=${this.mod}!${this.mode}`).then((j) => j.json()).catch(() => {
                alert.API();
            });

            this.stats = res;
        },

        async loadBest() {
            // prettier-ignore
            const best: Score[] = await fetch(`${config.api}/users/${this.id}/scores?m=${this.mod}!${this.mode}&b=1&p=${this.page.best++}`).then((j) => j.json()).catch(() => {
                alert.API();
            });

            const ids = this.best.map((i) => i.id);

            for (const score of best) {
                if (!ids.includes(score.id)) {
                    this.best.push(score);
                }
            }
        },

        async loadRecent() {
            // prettier-ignore
            const recent: Score[] = await fetch(`${config.api}/users/${this.id}/scores?m=${this.mod}!${this.mode}&b=0&p=${this.page.recent++}`).then((j) => j.json()).catch(() => {
                alert.API();
            });

            const ids = this.recent.map((i) => i.id);

            for (const score of recent) {
                if (!ids.includes(score.id)) {
                    this.recent.push(score);
                }
            }
        },

        async loadPinned() {
            // prettier-ignore
            const pinned: Score[] = await fetch(`${config.api}/users/${this.id}/pinned?m=${this.mod}!${this.mode}&p=${this.page.pinned++}`).then((j) => j.json()).catch(() => {
                alert.API();
            });

            const ids = this.pinned.map((i) => i.id);

            for (const score of pinned) {
                if (!ids.includes(score.id)) {
                    this.pinned.push(score);
                }
            }
        },

        updateEverything() {
            this.page.best = 0;
            this.page.recent = 0;
            this.page.pinned = 0;

            this.best = [];
            this.recent = [];

            this.loadStats();
            this.loadBest();
            this.loadRecent();
            this.loadPinned();
        },

        // ui
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

            router.push({ query: { g: n, m: this.mode } });
            this.updateEverything();
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

            router.push({ query: { g: this.mod, m: n } });
            this.updateEverything();
        },
    },
});
</script>

<style scoped>
.switch-btn {
    display: block !important;
}

.my-1 button {
    margin: 0 5px;
}
</style>
