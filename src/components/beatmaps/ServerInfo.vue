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
        <button
            class="button favourite"
            v-if="$store.getters.loggedIn"
            @click="favourite()"
        >
            {{ faved ? "Unfavourite" : "Favourite" }}
        </button>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Swal from "sweetalert2";

import * as alert from "../../util/error";
import config from "../../../config.json";

export default defineComponent({
    props: ["map"],
    data() {
        return {
            playing: false,
            audio: new Audio(),
            faved: false,
        };
    },

    async mounted() {
        if (!this.$store.getters.loggedIn) return;

        const faved: boolean = await fetch(
            `${config.api}/@me/fav?m=${this.map.set_id}`,
            {
                headers: {
                    Token: this.$store.state.token,
                },
            }
        )
            .then((f) => f.json())
            .catch(() => {
                alert.API();
            });

        if (typeof faved == "boolean") {
            this.faved = faved;
        }
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

        async favourite() {
            const res = await fetch(
                `${config.api}/@me/fav?m=${this.map.set_id}`,
                {
                    method: this.faved ? "DELETE" : "POST",
                    headers: {
                        Token: this.$store.state.token,
                    },
                }
            )
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            if (res?.code != 200) {
                Swal.fire({
                    title: "Invalid request!",
                    text: res.message,
                    icon: "error",
                });

                return;
            }

            Swal.fire({
                title: "Success!",
                icon: "success",
            });

            this.faved = !this.faved;
        },
    },

    watch: {
        $route() {
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

.favourite {
    background-color: #ffa600;
}

.button:active {
    opacity: 0.8;
}
</style>
