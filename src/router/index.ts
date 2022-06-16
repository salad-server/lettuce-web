import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Index from "../views/Index.vue";
import NotFound from "../views/NotFound.vue";

import beatmaps from "./routes/beatmaps";
import index from "./routes/index";
import info from "./routes/info";
import profiles from "./routes/profiles";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Index",
        component: Index,
    },

    ...beatmaps,
    ...index,
    ...info,
    ...profiles,

    {
        path: "/:notFound(.*)",
        component: NotFound,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
