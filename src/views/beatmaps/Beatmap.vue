<template>
    <Loading v-if="loading" />
    <div v-else-if="!error">
        <div class="columns is-centered mt-2">
            <div class="column is-10 has-text-centered" :style="bg">
                <h2 class="is-size-3">
                    {{ currentMap.artist }} - {{ currentMap.title }} [{{
                        currentMap.version
                    }}]
                </h2>
            </div>
        </div>

        <div class="columns is-centered">
            <div class="column is-9 is-centered drops">
                <InfoDrop :id="currentMap.id" :sets="setData" />
            </div>
        </div>

        <div class="columns is-centered">
            <div class="column is-3">
                <ServerInfo :map="currentMap" />
            </div>
            <div class="column is-6">
                <Info :map="currentMap" />
            </div>
        </div>

        <div class="columns is-centered">
            <Leaderboard
                v-if="allowlb"
                :id="currentMap.id"
                :play_mode="currentMap.mode"
            />
            <h3 v-else>No leaderboard available for this map!</h3>
        </div>
    </div>
    <Error v-else :msg="errorMsg" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Beatmap } from "@/types/beatmap";

import Loading from "@/components/global/Loading.vue";
import Error from "@/components/global/Error.vue";
import ServerInfo from "@/components/beatmaps/ServerInfo.vue";
import Info from "@/components/beatmaps/Info.vue";
import InfoDrop from "@/components/beatmaps/InfoDrop.vue";
import Leaderboard from "../../components/beatmaps/Leaderboard.vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["id"],
    components: { Loading, Error, ServerInfo, Info, InfoDrop, Leaderboard },

    data() {
        return {
            setData: [] as Beatmap[],
            currentMap: {} as Beatmap,

            loading: true,
            error: false,
            errorMsg: "",
        };
    },

    async mounted() {
        await this.loadMaps();
        this.loading = false;
    },

    methods: {
        async loadMaps() {
            // prettier-ignore
            const res = await fetch(`${config.api}/beatmap/${this.id}/sets`).then((j) => j.json()).catch(() => {
                this.error = true;
                this.errorMsg = "Could not load map!";
            });

            if (!res) return;
            if (res?.code == 400 || res.length <= 0) {
                this.error = true;
                this.errorMsg = "Beatmap not found!";
                return;
            }

            this.currentMap = res.find((r: Beatmap) => r.id == this.id);
            this.setData = res.sort(
                (a: Beatmap, b: Beatmap) => a.diff - b.diff
            );
        },
    },

    computed: {
        allowlb(): boolean {
            if (["Unranked", "Pending"].includes(this.currentMap.status)) {
                return config.leaderboard.allow_unranked;
            } else if (this.currentMap.status == "Loved") {
                return config.leaderboard.allow_loved;
            }

            return true;
        },

        bg() {
            return `
                background-image: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.6)), url("https://assets.ppy.sh/beatmaps/${this.currentMap.set_id}/covers/cover.jpg");
            `;
        },
    },
});
</script>

<style scoped>
.column {
    background-color: #eee;
}

.is-10 {
    color: #fff;
    padding: 100px 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
}

.drops {
    text-align: center;
}
</style>
