<template>
    <div>
        <h1>Home</h1>

        <p>
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Asperiores
            blanditiis similique quaerat. Voluptates sit necessitatibus illum
            voluptate aliquam cupiditate nihil.
        </p>
        <Records />
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Records from "@/components/index/Records.vue";

export default defineComponent({
    components: { Records },
});
</script>
