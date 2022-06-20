<template>
    <div>
        <table class="table">
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

        <div class="ranks has-text-centered">
            <div>
                <span class="rank rank-xh">XH</span>
                <span>{{ cstats.count_xh }}</span>
            </div>
            <div>
                <span class="rank rank-sh">SH</span>
                <span>{{ cstats.count_sh }}</span>
            </div>
            <br />
            <div>
                <span class="rank rank-x">X</span>
                <span>{{ cstats.count_x }}</span>
            </div>
            <div>
                <span class="rank rank-s">S</span>
                <span>{{ cstats.count_s }}</span>
            </div>
            <div>
                <span class="rank rank-a">A</span>
                <span>{{ cstats.count_a }}</span>
            </div>
        </div>
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

<style scoped>
.ranks {
    font-size: 2em;
    font-family: "Rubik", sans-serif;
    font-weight: lighter;
}

.ranks div {
    margin: 5px;
    display: inline-block;
}

.ranks span {
    margin-left: 5px;
}

.rank:after {
    content: ":";
    font-weight: lighter;
}

.table {
    width: 100%;
}

@media screen and (max-width: 768px) {
    table {
        text-align: center;
    }
}
</style>
