<template>
    <span v-for="set in sets" :key="set.id">
        <a
            :href="'/beatmaps/' + set.id"
            :class="id == set.id ? 'selected' : ''"
            class="tip"
        >
            <img
                :src="convertMode(set.mode)"
                :style="convertColor(set.diff.toFixed(2))"
                alt="Gamemode"
            />

            <div class="txt">{{ title(set) }}</div>
        </a>
    </span>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Beatmap } from "@/types/beatmap";

export default defineComponent({
    props: ["id", "sets"],
    methods: {
        convertColor(sr: number) {
            // prettier-ignore
            const colors = {
                "filter: invert(52%) sepia(91%) saturate(404%) hue-rotate(33deg) brightness(100%) contrast(88%);": [0.0, 1.99],
                "filter: invert(79%) sepia(7%) saturate(1264%) hue-rotate(142deg) brightness(103%) contrast(89%);": [2.0, 2.69],
                "filter: invert(84%) sepia(51%) saturate(3477%) hue-rotate(352deg) brightness(109%) contrast(74%);": [2.7, 3.99],
                "filter: invert(60%) sepia(72%) saturate(2255%) hue-rotate(297deg) brightness(99%) contrast(86%);": [4.0, 5.29],
                "filter: invert(39%) sepia(10%) saturate(2757%) hue-rotate(210deg) brightness(103%) contrast(84%);": [5.3, 6.49],
                "filter: invert(0%) sepia(1%) saturate(14%) hue-rotate(45deg) brightness(88%) contrast(97%);": [6.5, Infinity],
            };

            type c = keyof typeof colors;
            let src = "";

            for (let i in colors) {
                if (sr >= colors[i as c][0] && sr <= colors[i as c][1]) {
                    src = i;
                }
            }

            return src;
        },

        convertMode(mode: string) {
            return `/img/games/${mode.split("!")[1]}.png`;
        },

        title(set: Beatmap) {
            return `${set.version} (${set.diff.toFixed(2)}*)`;
        },
    },
});
</script>

<style scoped>
img {
    width: 35px;
}

a {
    display: inline-block;
    height: 50px;
    width: 50px;
    vertical-align: middle;
    margin: 0 5px;
    transition: all 0.1s ease-in-out;
}

a:hover,
.selected {
    transform: scale(1.3);
}

.tip {
    display: inline-block;
    position: relative;
    transform: scale(1);
}

.tip .txt {
    position: absolute;
    visibility: hidden;
    background-color: #000;
    width: 150px;
    color: #fff;
    text-align: center;
    border-radius: 4px;
    padding: 5px 3px;
    font-size: 0.8em;
    z-index: 1;
}

.tip:hover .txt {
    visibility: visible;
}
</style>
