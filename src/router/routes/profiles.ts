import Profile from "@/views/profiles/Profile.vue";
import Settings from "@/views/profiles/Settings.vue";
import Picture from "@/views/profiles/Picture.vue";

export default [
    {
        path: "/profile",
        name: "Profile Settings",
        component: Settings,
    },
    {
        path: "/profile/picture",
        name: "Profile Picture",
        component: Picture,
    },
    {
        path: "/profile/:id",
        name: "Profile Info",
        props: true,
        component: Profile,
    },
];
