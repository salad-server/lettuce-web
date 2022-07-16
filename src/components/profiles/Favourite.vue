<template>
    <div>
        <div class="columns is-mobile is-multiline">
            <FavItem
                v-for="fav in data"
                :data="fav"
                :key="fav.set_id"
                :playing="playing[fav.set_id]"
                @play="togglePlay(fav.set_id)"
            />
        </div>

        <button class="button" @click="$emit('clicked')">Load More</button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import FavItem from "./items/FavMaps.vue";

export default defineComponent({
    props: ["data"],
    components: { FavItem },
    data() {
        return {
            audio: new Audio(),
            playing: {} as {
                [id: number]: boolean;
            },
        };
    },

    mounted() {
        this.resetAudio();
    },

    methods: {
        resetAudio() {
            for (const d of this.data) {
                this.playing[d.set_id] = false;
            }
        },

        togglePlay(id: number) {
            for (const d of this.data) {
                if (d.set_id != id) {
                    this.playing[d.set_id] = false;
                }
            }

            this.audio.pause();
            this.playing[id] = !this.playing[id];

            if (this.playing[id]) {
                this.audio = new Audio(`https://b.ppy.sh/preview/${id}.mp3`);
                this.audio.play();

                return;
            }
        },
    },

    watch: {
        $route() {
            this.audio.pause();
            this.resetAudio();
        },
    },
});
</script>

<style scoped>
.button {
    display: block;
    margin: 0 auto;
}
</style>
