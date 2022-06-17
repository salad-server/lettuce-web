<template>
    <Loading v-if="loading" />
    <div v-else-if="!error">
        <Map :beatmap="score.map" />
        <hr />
        <Score :score="score" />
    </div>
    <Error v-else :msg="errorMsg" />
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
            });

            if (!res) return;
            if (res.id == 0 || [400, 500].includes(res?.code || 0)) {
                this.error = true;
                this.errorMsg = "Score could not be located! Check ID?";
            }

            this.score = res;
        },
    },
});
</script>
