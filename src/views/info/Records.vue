<template>
    <Loading v-if="!loaded" />
    <div v-else class="columns is-mobile is-centered mt-5">
        <div class="column is-8">
            <h3 class="subtitle is-size-3">{{ show }} Records</h3>
            <p class="my-5">
                A list of pp/score records set on this server. You can sort
                between the two using the buttons below. Please note that you
                can only set a record on a ranked/approved map. Loved and
                unranked maps do not count.
            </p>

            <button
                class="button my-3"
                :class="isPP ? 'is-primary' : ''"
                @click="show = 'PP'"
            >
                PP
            </button>
            <button
                class="button my-3"
                :class="!isPP ? 'is-primary' : ''"
                @click="show = 'Score'"
            >
                Score
            </button>

            <div class="grid-wrapper" v-if="show == 'PP'">
                <RecordItem
                    v-for="record in records"
                    :key="record.pp.id"
                    :score="record.pp"
                />
            </div>
            <div class="grid-wrapper" v-else>
                <RecordItem
                    v-for="record in records"
                    :key="record.score.id"
                    :score="record.score"
                />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Records } from "@/types/stats";
import * as alert from "@/util/error";

import Loading from "@/components/global/Loading.vue";
import RecordItem from "@/components/info/RecordItem.vue";
import config from "../../../config.json";

export default defineComponent({
    components: { Loading, RecordItem },
    data() {
        return {
            show: "PP",
            loaded: false,
            records: {} as Records,
        };
    },

    async mounted() {
        await this.loadRecords();
        this.loaded = true;
    },

    methods: {
        async loadRecords() {
            const res: Records = await fetch(`${config.api}/score`)
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            this.records = res;
        },
    },

    computed: {
        isPP(): boolean {
            return this.show == "PP";
        },
    },
});
</script>

<style scoped>
.grid-wrapper {
    display: grid;
    grid-template-areas: "auto auto";
    gap: 60px;
}

@media screen and (max-width: 1350px) {
    .grid-wrapper {
        grid-template-areas: "auto";
    }
}
</style>
