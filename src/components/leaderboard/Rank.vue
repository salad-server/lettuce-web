<template>
    <tr>
        <td>#{{ page + 1 }}</td>
        <td>
            <router-link :to="pflink">
                {{ rank.user.country }} | {{ rank.user.username }}
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

export default defineComponent({
    props: ["page", "rank", "mode"],
    computed: {
        pflink() {
            const m = this.mode.split("!");

            return `/profile/${this.rank.user.id}?g=${m[0]}&m=${m[1]}`;
        },
    },

    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },
});
</script>
