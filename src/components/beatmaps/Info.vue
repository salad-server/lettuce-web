<template>
    <table class="table">
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
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";

export default defineComponent({
    props: ["map"],
    computed: {
        len() {
            return moment.utc(this.map.total_length * 1000).format("HH:mm:ss");
        },
    },
});
</script>

<style scoped>
table {
    width: 100%;
    overflow-x: auto;
}

@media screen and (max-width: 768px) {
    table {
        width: 95%;
        margin: auto;
    }
}
</style>
