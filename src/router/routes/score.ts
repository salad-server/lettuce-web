import Records from "@/views/score/Records.vue";
import Score from "@/views/score/Score.vue";

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
