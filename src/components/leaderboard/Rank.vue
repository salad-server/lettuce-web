<template>
    <tr>
        <td class="td-rank">#{{ page + 1 }}</td>
        <td class="player">
            <router-link :to="pflink">
                <img :src="pfp" alt="Profile Picture" width="40" height="40" />
                <img :src="flag" :alt="rank.user.country" width="20" />
                <span>
                    {{ rank.user.username }}
                </span>
            </router-link>
        </td>
        <td>{{ comma(rank.pp) }}pp</td>
        <td>{{ comma(rank.total_score) }}</td>
        <td>{{ comma(rank.ranked_score) }}</td>
        <td>{{ comma(rank.playcount) }}</td>
    </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["page", "rank", "mode"],
    computed: {
        pfp() {
            return `${config.avatar}/${this.rank.user.id}`;
        },

        pflink() {
            const m = this.mode.split("!");

            return `/profile/${this.rank.user.id}?g=${m[0]}&m=${m[1]}`;
        },

        flag() {
            return `/img/flags/flag-${this.rank.user.country}.svg`;
        },
    },

    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },
});
</script>
