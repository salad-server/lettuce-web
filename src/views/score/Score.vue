<template>
    <Loading v-if="loading" />
    <div v-else-if="!error">
        <div class="columns is-centered mt-2">
            <div class="column is-10" :style="bg">
                <Score :score="score" />
            </div>
        </div>
        <div class="columns is-centered mt-2">
            <div class="column is-9">
                <Map :beatmap="score.map" :mode="score.play_mode" />
            </div>
        </div>
    </div>
    <Error v-else :code="errCode" :msg="errorMsg" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { AdvancedScore } from "@/types/scores";

import Loading from "@/components/global/Loading.vue";
import Error from "@/components/global/Error.vue";
import Map from "@/components/score/Map.vue";
import Score from "@/components/score/Score.vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["id"],
    components: { Loading, Error, Map, Score },
    data() {
        return {
            loading: true,
            score: {} as AdvancedScore,
            error: false,
            errorMsg: "",
            errCode: 0,
        };
    },

    async mounted() {
        await this.loadScore();

        this.loading = false;
    },

    methods: {
        async loadScore() {
            // prettier-ignore
            const res: AdvancedScore = await fetch(`${config.api}/score/${this.id}`).then(j => j.json()).catch(() => {
                this.error = true;
                this.errorMsg = "Error loading score!";
                this.errCode = 500;
            });

            if (!res) return;
            if (res.id == 0 || [400, 500].includes(res?.code || 0)) {
                this.error = true;
                this.errorMsg = "Score could not be located! Check ID?";
                this.errCode = 400;
            }

            this.score = res;
        },
    },

    computed: {
        bg() {
            return `
                background-image: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.6)), url("https://assets.ppy.sh/beatmaps/${this.score.map.set_id}/covers/cover.jpg");
            `;
        },
    },
});
</script>

<style scoped>
.is-10 {
    padding: 40px;
    color: #fff;
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
}

.is-9 {
    background-color: #eee;
}
</style>
