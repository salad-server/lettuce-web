<template>
    <div class="column is-half-desktop is-half-tablet is-full-mobile">
        <div class="card">
            <div class="card-image">
                <figure class="image">
                    <img :src="cover" class="cover" />
                </figure>
                <div class="card-overlay">
                    <p :class="status">{{ data.children[0].status }}</p>
                    <div id="play">
                        <img
                            width="25"
                            :src="playPause"
                            @click="$emit('play')"
                        />
                    </div>
                    <div class="infos">
                        <router-link
                            :title="data.title"
                            :to="`/beatmaps/${data.children[0].id}`"
                        >
                            {{ data.title.replace(/(.{35})..+/, "$1…") }}
                        </router-link>
                        <p>
                            By <b>{{ data.artist }}</b>
                        </p>
                    </div>
                </div>
            </div>
            <div class="card-content">
                <div class="content info-drop-card">
                    <InfoDrop id="0" :sets="data.children" />
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import InfoDrop from "@/components/beatmaps/InfoDrop.vue";

export default defineComponent({
    props: ["data", "playing"],
    components: { InfoDrop },
    computed: {
        cover(): string {
            return `https://assets.ppy.sh/beatmaps/${this.data.set_id}/covers/card.jpg`;
        },

        status(): string {
            return `status status-${this.data.children[0].status.toLowerCase()}`;
        },

        playPause(): string {
            return `/img/icons/${this.playing ? "pause" : "play"}.svg`;
        },
    },
});
</script>

<style scoped>
.card-overlay {
    display: block;
    position: absolute;
    top: 0;
    margin: 2%;
    width: 96%;
}

.card-image {
    background-color: #000;
}

.cover {
    opacity: 0.5;
}

.status {
    background-color: rgba(0, 0, 0, 0.5);
    display: inline-block;
    border-radius: 5px;
    padding: 5px 10px;
    font-weight: bolder;
    font-size: 1vw;
}

.infos * {
    color: #fff !important;
}

.infos a {
    font-size: 1.5vw;
}

.infos a:hover {
    text-decoration: underline;
    opacity: 0.8;
}

.infos p {
    font-size: 1.1vw;
}
.image {
    overflow: hidden;
    max-height: 150px !important;
}

#play {
    position: absolute;
    left: 90%;
    top: 10%;
    filter: invert(100%) sepia(21%) saturate(2%) hue-rotate(277deg)
        brightness(111%) contrast(100%);
}

@media screen and (min-width: 2100px) {
    .infos a {
        font-size: 1.2em !important;
    }

    .status,
    .infos p {
        font-size: 1em !important;
    }
}

@media screen and (max-width: 768px) {
    .status {
        font-size: 1em;
    }

    .infos a {
        margin-top: 20px;
        font-size: 1.5em;
    }

    .infos p {
        font-size: 1.1em;
    }
}
</style>

<style>
.info-drop-card .tip,
.info-drop-card span img {
    width: 2em;
    height: 2em;
}
</style>
