<template>
    <div>
        <h1>Documentation</h1>

        <Loading v-if="!loaded" />
        <ul v-else>
            <li v-for="doc in docs" :key="doc">
                <router-link :to="'/info/' + doc">
                    {{ clean(doc) }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import * as alert from "@/util/error";
import Loading from "@/components/global/Loading.vue";
import config from "../../../config.json";

export default defineComponent({
    components: { Loading },
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
    },
});
</script>
