<template>
    <div class="columns is-mobile is-centered mt-2">
        <div class="column dcard is-2">
            <h1 class="is-size-2 title">Documentation</h1>
            <p class="subtitle is-size-5">
                Here you can find documentation about the server!
            </p>

            <LoadingMini v-if="!loaded" />
            <ul v-else class="content">
                <li v-for="doc in docs" :key="doc">
                    <router-link :to="'/info/' + url(doc)">
                        {{ clean(doc) }}
                    </router-link>
                </li>

                <img src="/img/errors/i-forgor.jpg" alt="I require support" />
            </ul>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import * as alert from "@/util/error";
import LoadingMini from "@/components/global/LoadingMini.vue";
import config from "../../../config.json";

export default defineComponent({
    components: { LoadingMini },
    data() {
        return {
            loaded: false,
            docs: [],
        };
    },

    async mounted() {
        this.docs = await fetch(`${config.api}/docs`)
            .then((j) => j.json())
            .catch(() => alert.API());

        this.loaded = true;
    },

    methods: {
        clean(s: string) {
            return s.charAt(0).toUpperCase() + s.slice(1, s.length - 3);
        },

        url(doc: string) {
            return doc.slice(0, doc.length - 3);
        },
    },
});
</script>

<style scoped>
.dcard {
    width: 70% !important;
}

img {
    float: right;
}

ul {
    list-style-type: disc;
    margin-left: 1em;
}
</style>
