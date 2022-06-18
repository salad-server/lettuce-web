<template>
    <Loading v-if="loading" />
    <Error v-else-if="error" :msg="errorMsg" />
    <div v-html="tohtml"></div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { createMarkdown } from "safe-marked";

const markdown = createMarkdown();

import Loading from "@/components/global/Loading.vue";
import Error from "@/components/global/Error.vue";
import config from "../../../config.json";

export default defineComponent({
    props: ["doc"],
    components: { Loading, Error },
    data() {
        return {
            loading: true,
            error: false,
            errorMsg: "",
            md: "",
        };
    },

    async mounted() {
        await fetch(`${config.api}/docs/${this.doc}`)
            .then(async (t) => {
                if (t.status != 200) {
                    throw t.status;
                }

                this.md = await t.text();
            })
            .catch(() => {
                this.error = true;
                this.errorMsg = "Documentation not found!";
            });

        this.loading = false;
    },

    computed: {
        tohtml(): string {
            return markdown(this.md);
        },
    },
});
</script>
