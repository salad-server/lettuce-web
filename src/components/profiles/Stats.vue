<template>
    <div>
        <table>
            <tr>
                <td>Total Score</td>
                <td>{{ cstats.total_score }}</td>
            </tr>
            <tr>
                <td>Ranked Score</td>
                <td>{{ cstats.ranked_score }}</td>
            </tr>
            <tr>
                <td>PP</td>
                <td>{{ cstats.pp }}pp</td>
            </tr>
            <tr>
                <td>Playcount</td>
                <td>{{ cstats.playcount }}</td>
            </tr>
            <tr>
                <td>Accuracy</td>
                <td>{{ acc }}%</td>
            </tr>
            <tr>
                <td>Max Combo</td>
                <td>{{ cstats.max_combo }}</td>
            </tr>
            <tr>
                <td>Total Hits</td>
                <td>{{ cstats.total_hits }}</td>
            </tr>
            <tr>
                <td>Replays Watched</td>
                <td>{{ cstats.watched_replays }}</td>
            </tr>
        </table>

        <table>
            <tr>
                <td>XH</td>
                <td>{{ cstats.count_xh }}</td>
            </tr>
            <tr>
                <td>SH</td>
                <td>{{ cstats.count_sh }}</td>
            </tr>
            <tr>
                <td>X</td>
                <td>{{ cstats.count_x }}</td>
            </tr>
            <tr>
                <td>S</td>
                <td>{{ cstats.count_s }}</td>
            </tr>
            <tr>
                <td>A</td>
                <td>{{ cstats.count_a }}</td>
            </tr>
        </table>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

// https://stackoverflow.com/a/62256954
const comma = (n: number) => n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");

export default defineComponent({
    props: ["stats"],
    data() {
        // eslint-disable-next-line
        return { cstats: {} as any };
    },

    computed: {
        acc() {
            return (this.stats.acc || 0).toFixed(2);
        },
    },

    watch: {
        stats() {
            for (const i of Object.keys(this.stats)) {
                if (this.stats[i] != undefined) {
                    this.cstats[i] = comma(this.stats[i]);
                }
            }
        },
    },
});
</script>
