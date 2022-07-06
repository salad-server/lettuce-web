import Leaderboard from "@/views/index/Leaderboard.vue";
import Login from "@/views/index/Login.vue";

export default [
    {
        path: "/leaderboard",
        name: "Leaderboard",
        component: Leaderboard,
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
];
