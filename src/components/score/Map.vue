<template>
    <h2>
        <router-link :to="'/beatmaps/' + beatmap.map_id">
            {{ beatmap.artist }} - {{ beatmap.title }} [{{ beatmap.version }}]
        </router-link>
    </h2>
    <h4>
        Mapped by
        <a :href="'https://osu.ppy.sh/u/' + beatmap.creator" target="_blank">{{
            beatmap.creator
        }}</a>
    </h4>
    <table>
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

// TODO: Download replay

export default defineComponent({
    props: ["beatmap"],
    computed: {
        len() {
            return moment.utc(this.beatmap.len * 1000).format("HH:mm:ss");
        },
    },
});
</script>
