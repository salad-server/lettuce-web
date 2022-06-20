<template>
    <div class="inner">
        <img src="https://a.ppy.sh/" />
        <h4>
            Mapped by
            <a :href="creator" target="_blank">{{ map.creator }}</a>
        </h4>
    </div>

    <div class="inner-data">
        <div>
            <b>Difficulty</b>
            <span>{{ map.diff.toFixed(2) }}*</span>
        </div>
        <div>
            <b>Ranked</b>
            <span :class="status">{{ map.status }}</span>
        </div>
        <div>
            <b>Mode</b>
            <span>{{ map.mode }}</span>
        </div>
    </div>

    <div class="outer-buttons">
        <!-- TODO: Play button -->
        <button class="button play" @click="togglePlay()">
            {{ playing ? "Stop" : "Play" }}
        </button>
        <button class="button download" @click="open('osu://dl/' + map.set_id)">
            Download
        </button>
        <button
            class="button bancho"
            @click="open('https://osu.ppy.sh/b/' + map.id)"
        >
            Bancho
        </button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["map"],
    data() {
        return {
            playing: false,
            audio: new Audio(),
        };
    },

    computed: {
        creator() {
            return "https://osu.ppy.sh/u/" + this.map.creator;
        },

        status() {
            return `status-${this.map.status.toLowerCase()}`;
        },
    },

    methods: {
        open(url: string) {
            window.location.href = url;
        },

        togglePlay() {
            this.playing = !this.playing;

            if (this.playing) {
                this.audio = new Audio(
                    `https://b.ppy.sh/preview/${this.map.set_id}.mp3`
                );

                this.audio.play();
                return;
            }

            this.audio.pause();
        },
    },
});
</script>

<style scoped>
.inner {
    text-align: center;
}

.inner * {
    display: inline-block;
    vertical-align: middle;
}

.inner img {
    margin-right: 1em;
    width: 50px;
    height: 50px;
}

.inner-data {
    line-height: 2em;
    margin-top: 5%;
    text-align: center;
}

.inner-data b:after {
    content: ": ";
}

.button {
    margin-bottom: 3px;
}

.outer-buttons {
    display: flex;
    width: 80%;
    margin: auto;
    justify-content: space-around;
    flex-direction: column;
}

.button {
    font-weight: bolder;
    color: #fff;
    border-radius: 8px;
}

.bancho {
    background-color: #e61a8d;
}

.download {
    background-color: #3986ac;
}

.play {
    background-color: #9400ff;
}

.button:active {
    opacity: 0.8;
}
</style>
