<template>
    <Loading v-if="loading" />
    <Error v-else-if="error" :msg="errorMsg" />
    <div class="columns is-mobile is-centered mt-2">
        <div class="column dcard is-2">
            <div class="content" v-html="tohtml"></div>
        </div>
    </div>
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
        await fetch(`${config.api}/docs/${this.doc}.md`)
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

<style scoped>
.dcard {
    width: 70% !important;
}
</style>
