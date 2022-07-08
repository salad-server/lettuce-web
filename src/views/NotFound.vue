<template>
    <Error code="404" msg="Page not found!" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Error from "@/components/global/Error.vue";

export default defineComponent({
    components: { Error },
});
</script>
