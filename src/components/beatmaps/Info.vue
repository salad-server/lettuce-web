<template>
    <div>
        <img :src="cover" alt="Cover" />
        <h2>{{ map.artist }} - {{ map.title }} [{{ map.version }}]</h2>
        <h4>
            Mapped by
            <a :href="creator" target="_blank">{{ map.creator }}</a>
        </h4>
        <h4></h4>
    </div>

    <div>
        <table>
            <tr>
                <td>Plays/Passes</td>
                <td>{{ map.plays }}/{{ map.passes }}</td>
            </tr>
            <tr>
                <td>Length</td>
                <td>{{ len }}</td>
            </tr>
            <tr>
                <td>Max Combo</td>
                <td>{{ map.max_combo }}</td>
            </tr>
            <tr>
                <td>Difficulty</td>
                <td>{{ map.diff }}</td>
            </tr>
            <tr>
                <td>Ranked</td>
                <td>{{ map.status }}</td>
            </tr>
            <tr>
                <td>Mode</td>
                <td>{{ map.mode }}</td>
            </tr>
            <tr>
                <td>OD</td>
                <td>{{ map.od }}</td>
            </tr>
            <tr>
                <td>HP</td>
                <td>{{ map.hp }}</td>
            </tr>
            <tr>
                <td>CS</td>
                <td>{{ map.cs }}</td>
            </tr>
            <tr>
                <td>AR</td>
                <td>{{ map.ar }}</td>
            </tr>
        </table>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";

export default defineComponent({
    props: ["map"],
    computed: {
        cover() {
            return `https://assets.ppy.sh/beatmaps/${this.map.set_id}/covers/cover.jpg`;
        },

        creator() {
            return "https://osu.ppy.sh/u/" + this.map.creator;
        },

        len() {
            return moment.utc(this.map.total_length * 1000).format("HH:mm:ss");
        },
    },
});
</script>
