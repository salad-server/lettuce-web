import Profile from "@/views/profiles/Profile.vue";
import Settings from "@/views/profiles/Settings.vue";
import Score from "@/views/profiles/Score.vue";

export default [
    {
        path: "/profile",
        name: "Profile Settings",
        component: Settings,
    },
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
];
