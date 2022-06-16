<template>
    <tr>
        <td>
            <h1>{{ score.rank }}</h1>
        </td>
        <td>
            <router-link :to="'/beatmaps/' + score.map.map_id">
                <b>{{ score.map.title }} [{{ score.map.version }}]</b>
            </router-link>
            <div>
                {{ comma(score.c300) }} / {{ comma(score.c100) }} /
                {{ comma(score.c50) }} / {{ comma(score.miss) }} |
                {{ comma(score.combo) }}x
            </div>
            <div>
                <b>{{ comma(score.score) }}</b>
            </div>
        </td>
        <td>
            <div v-html="mods"></div>
        </td>
        <td>
            <h3>{{ score.pp.toFixed(2) }}pp</h3>
            <router-link :to="`/score/${score.id}`">Info</router-link> |
            <!-- prettier-ignore -->
            <router-link :to="`/score/${score.id}/download`">Download</router-link>
            <div v-if="score.submission_status != 2">unsubmitted</div>
        </td>
    </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import * as m from "@/util/mods";

export default defineComponent({
    props: ["score", "mode"],
    methods: {
        comma: (n: number) =>
            n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ","),
    },

    computed: {
        mods() {
            const mode = ["std", "taiko", "catch", "mania"].indexOf(this.mode);
            const mods = m.convertMods(this.score.mods, mode);
            const str = m.modstr(mods);

            // prettier-ignore
            return str.map((s) => `<img src="/img/${s.toLowerCase()}.png" alt="${s}" />`).join("");
        },
    },
});
</script>
