<template>
    <div>
        <img :src="pfp" alt="Profile Picture" />
        <h1>{{ info.username }}</h1>
        <h3>{{ info.country }}</h3>
        <h3>{{ fixTime(info.register_time) }}</h3>
        <h3>{{ fixTime(info.last_online) }}</h3>

        <!-- TODO: Badges -->
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import config from "../../../config.json";
import moment from "moment";

export default defineComponent({
    props: ["info", "id"],
    computed: {
        pfp() {
            return `${config.avatar}/${this.id}`;
        },
    },

    methods: {
        fixTime(t: number) {
            return moment(new Date(t * 1000)).fromNow();
        },
    },
});
</script>
