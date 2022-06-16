import About from "@/views/info/About.vue";
import Stats from "@/views/info/Stats.vue";

export default [
    {
        path: "/info/about",
        name: "About Info",
        component: About,
    },
    {
        path: "/info/stats",
        name: "About Stats",
        component: Stats,
    },
];
