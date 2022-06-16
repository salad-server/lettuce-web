import Profile from "@/views/profiles/Profile.vue";
import Score from "@/views/profiles/Score.vue";

export default [
    {
        path: "/profile/:id",
        name: "Profile Info",
        props: true,
        component: Profile,
    },
    {
        path: "/score/:id",
        name: "Score Info",
        props: true,
        component: Score,
    },
    {
        path: "/score/:id/download",
        name: "Score Download",
        props: true,
        component: Score,
    },
];
