import Listing from "@/views/info/Listing.vue";
import Records from "@/views/info/Records.vue";
import Doc from "@/views/info/Doc.vue";

export default [
    {
        path: "/info",
        name: "Documentation",
        component: Listing,
    },
    {
        path: "/info/records",
        name: "Server Records",
        component: Records,
    },
    {
        path: "/info/:doc",
        name: "Document",
        props: true,
        component: Doc,
    },
];
