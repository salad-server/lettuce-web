<template>
    <div class="has-text-centered">
        <div class="middle">
            <div class="title is-primary">{{ code }}</div>
            <div class="oops">Request Error</div>
            <div>{{ msg }}</div>
        </div>
    </div>
</template>

<style scoped>
.has-text-centered {
    height: 99vh;
    width: 100%;
    position: absolute;
    top: 0;
    left: 0;
}

.title {
    font-size: 8em;
    line-height: 0.8em;
    color: hsl(171, 100%, 41%);
}

.middle {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
}

.oops {
    font-size: 2em;
}
</style>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: ["code", "msg"],
});
</script>
