import Listing from "@/views/info/Listing.vue";
import Doc from "@/views/info/Doc.vue";

export default [
    {
        path: "/info",
        name: "Documentation",
        component: Listing,
    },
    {
        path: "/info/:id",
        name: "Document",
        props: true,
        component: Doc,
    },
];
