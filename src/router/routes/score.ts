import Records from "@/views/info/Records.vue";
import Score from "@/views/profiles/Score.vue";

export default [
    {
        path: "/score",
        name: "Server Records",
        component: Records,
    },
    {
        path: "/score/:id",
        name: "Score Info",
        props: true,
        component: Score,
    },
];
