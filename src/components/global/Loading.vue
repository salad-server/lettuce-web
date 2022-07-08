<template>
    <div class="loader-wrapper">
        <div class="loader is-loading"></div>
    </div>
</template>

<style scoped>
.loader-wrapper {
    position: absolute;
    top: 0;
    height: 99vh;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.loader {
    width: 100px;
    height: 100px;
}
</style>
