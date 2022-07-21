<template>
    <div class="columns">
        <div class="column is-half">
            <img :src="pfp" alt="Profile Picture" height="150" width="150" />
        </div>
        <div class="column is-half">
            <h1 class="title is-3">
                <img :src="flag" :alt="info.country" width="25" />
                {{ info.username }}
            </h1>
            <h3>
                <img class="icon" src="/img/icons/join.svg" alt="Join icon" />
                <b>Joined</b>
                {{ fixTime(info.register_time) }}
            </h3>
            <h3>
                <img class="icon" src="/img/icons/login.svg" alt="Login icon" />
                <b>Last seen</b>
                {{ fixTime(info.last_online) }}
            </h3>

            <!-- TODO: Badges -->
        </div>
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
            return `${config.api}/users/${this.id}/avatar`;
        },

        flag() {
            return `/img/flags/flag-${this.info.country}.svg`;
        },
    },

    methods: {
        fixTime(t: number) {
            return moment(new Date(t * 1000)).fromNow();
        },
    },
});
</script>

<style scoped>
.icon {
    width: 1em;
    margin-right: 5px;
    vertical-align: middle;
}

@media screen and (max-width: 768px) {
    .is-half {
        text-align: center;
    }
}
</style>
