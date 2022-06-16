import Map from "@/views/beatmaps/Beatmap.vue";
import Listing from "@/views/beatmaps/Listing.vue";

export default [
    {
        path: "/beatmaps/:id",
        name: "Beatmap Leaderboard",
        props: true,
        component: Map,
    },
    {
        path: "/beatmaps/listing",
        name: "Beatmap Listing",
        component: Listing,
    },
];
