/* eslint-disable */
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import config from "../../config.json";

import Index from "../views/Index.vue";
import NotFound from "../views/NotFound.vue";

import beatmaps from "./routes/beatmaps";
import index from "./routes/index";
import info from "./routes/info";
import profiles from "./routes/profiles";
import score from "./routes/score";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component: Index,
    },

    ...beatmaps,
    ...index,
    ...info,
    ...profiles,
    ...score,

    {
        path: "/:notFound(.*)",
        name: "404 Page not found!",
        component: NotFound,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach((to: any, from, next) => {
    (document.title as any) = `${to.name} | ${config.name}`;
    next();
});

export default router;
