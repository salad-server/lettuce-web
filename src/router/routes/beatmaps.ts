import Map from "@/views/beatmaps/Beatmap.vue";

// TODO: Listing
// TODO: Search

export default [
    {
        path: "/beatmaps/:id",
        name: "Beatmap Leaderboard",
        props: true,
        component: Map,
    },
];
