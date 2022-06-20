<template>
    <Navbar />
    <router-view />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Navbar from "@/components/global/Navbar.vue";

// TODO: records style
// TODO: something in home

import "@/assets/main.scss";

export default defineComponent({
    name: "HomeView",
    components: {
        Navbar,
    },
});
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Rubik:wght@500&display=swap");

/* prettier-ignore */
.rank-xh { color: #CDE7E7; } /* prettier-ignore */
.rank-sh { color: #CDE7E7; } /* prettier-ignore */
.rank-x  { color: #FFCC22; } /* prettier-ignore */
.rank-s  { color: #FFCC22; } /* prettier-ignore */
.rank-a  { color: #2BFF35; } /* prettier-ignore */
.rank-b  { color: #3D97FF; } /* prettier-ignore */
.rank-c  { color: #FF56dA; } /* prettier-ignore */
.rank-d  { color: #FF6262; } /* prettier-ignore */
.rank-f  { color: #FF5959; } /* prettier-ignore */
.rank {
    font-family: "Rubik", sans-serif;
    font-weight: bolder;
}

/* prettier-ignore */
.status-unsubmitted { color: rgba(255, 255, 255, 0.6); } /* prettier-ignore */
.status-pending { color: rgb(255, 126, 40); } /* prettier-ignore */
.status-loved { color: rgb(235, 104, 126); } /* prettier-ignore */
.status-qualified { color: rgb(117, 255, 104, 1); } /* prettier-ignore */

.td-rank {
    width: 0.1%;
    white-space: nowrap;
}

.dcard {
    box-shadow: 0 1px 1px #000;
}

.enabled {
    color: red;
}
</style>
