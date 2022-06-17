<template>
    <Loading v-if="loading" />
    <div v-else-if="!error">
        <div>
            <Info :map="currentMap" />
            <InfoDrop :id="currentMap.id" :sets="setData" />
        </div>

        <Leaderboard :id="currentMap.id" :play_mode="currentMap.mode" />
    </div>
    <Error v-else :msg="errorMsg" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Beatmap } from "@/types/beatmap";

import Loading from "@/components/global/Loading.vue";
import Error from "@/components/global/Error.vue";
import Info from "@/components/beatmaps/Info.vue";
import InfoDrop from "@/components/beatmaps/InfoDrop.vue";
import Leaderboard from "../../components/beatmaps/Leaderboard.vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["id"],
    components: { Loading, Error, Info, InfoDrop, Leaderboard },

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

            if (res?.error || res.length <= 0) {
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
});
</script>
