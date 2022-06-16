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
        path: "/profile/score/:id",
        name: "Profile Score",
        props: true,
        component: Score,
    },
];
