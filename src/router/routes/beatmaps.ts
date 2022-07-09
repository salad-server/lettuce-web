import Map from "@/views/beatmaps/Beatmap.vue";

export default [
    {
        path: "/beatmaps/:id",
        name: "Beatmap Leaderboard",
        props: true,
        component: Map,
    },
];
