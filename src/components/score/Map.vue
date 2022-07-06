<template>
    <div class="infos">
        <h1 class="title">Beatmap Info</h1>
        <h4 class="subtitle">
            Mapped by
            <a
                :href="'https://osu.ppy.sh/u/' + beatmap.creator"
                target="_blank"
                >{{ beatmap.creator }}</a
            >
        </h4>

        <img :src="icon" alt="Playmode Icon" />
    </div>

    <table class="table">
        <tr>
            <td>Beatmapset</td>
            <td>
                <a :href="'https://osu.ppy.sh/beatmapsets/' + beatmap.set_id">{{
                    beatmap.set_id
                }}</a>
            </td>
        </tr>
        <tr>
            <td>Beatmap</td>
            <td>
                <a :href="'https://osu.ppy.sh/beatmaps/' + beatmap.map_id">{{
                    beatmap.map_id
                }}</a>
            </td>
        </tr>
        <tr>
            <td>Length</td>
            <td>{{ len }}</td>
        </tr>
        <tr>
            <td>BPM</td>
            <td>{{ beatmap.bpm }}</td>
        </tr>
        <tr>
            <td>Difficulty</td>
            <td>{{ beatmap.diff.toFixed(2) }}</td>
        </tr>
        <tr>
            <td>Status</td>
            <td>{{ beatmap.status }}</td>
        </tr>
        <tr>
            <td>OD</td>
            <td>{{ beatmap.od }}</td>
        </tr>
        <tr>
            <td>CS</td>
            <td>{{ beatmap.cs }}</td>
        </tr>
        <tr>
            <td>AR</td>
            <td>{{ beatmap.ar }}</td>
        </tr>
        <tr>
            <td>HP</td>
            <td>{{ beatmap.hp }}</td>
        </tr>
    </table>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";

export default defineComponent({
    props: ["beatmap", "mode"],
    computed: {
        len() {
            return moment.utc(this.beatmap.len * 1000).format("HH:mm:ss");
        },

        icon() {
            return `/img/games/${this.mode.split("!")[1]}.png`;
        },
    },
});
</script>

<style scoped>
table {
    width: 100%;
}

.infos {
    position: relative;
}

img {
    width: 50px;
    height: 50px;
    position: absolute;
    top: 0;
    left: calc(100% - 80px);
    transform: rotate(-10deg);
}
</style>
