<template>
    <Loading v-if="!loaded" />
    <div v-else>
        <h3>{{ show }} Records</h3>

        <div v-if="show == 'PP'">
            <RecordItem
                v-for="record in records"
                :key="record.pp.id"
                :score="record.pp"
            />
        </div>
        <div v-else>
            <RecordItem
                v-for="record in records"
                :key="record.score.id"
                :score="record.score"
            />
        </div>

        <button :class="isPP ? 'enabled' : ''" @click="show = 'PP'">PP</button>
        <button :class="!isPP ? 'enabled' : ''" @click="show = 'Score'">
            Score
        </button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Records } from "@/types/stats";
import Swal from "sweetalert2";

import Loading from "../global/Loading.vue";
import config from "../../../config.json";
import RecordItem from "./RecordItem.vue";

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
                    Swal.fire({
                        title: "API Error!",
                        text: "Check your connection. Please report this to a staff member if the problem persists.",
                        icon: "error",
                    });
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
