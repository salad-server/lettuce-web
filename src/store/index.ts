import { createStore } from "vuex";
import { AuthInfo, AuthUser } from "@/types/user";

// TODO: Auto logout

export default createStore({
    state: {
        token: "",
        username: "",
        expire: 0,
        priv: 0,
        id: 0,
    },

    mutations: {
        login(state, token) {
            const parsed: AuthUser = JSON.parse(atob(token.split(".")[1]));
            parsed.user = JSON.parse(String(parsed.user)) as AuthInfo;

            // save in state
            state.username = parsed.user.username;
            state.id = parsed.user.id;
            state.priv = parsed.user.priv;
            state.expire = parsed.exp;
            state.token = token;

            localStorage.setItem("token", token);
        },

        logout(state) {
            state.token = "";
            state.username = "";
            state.expire = 0;
            state.priv = 0;
            state.id = 0;

            localStorage.clear();
            window.location.href = "/";
        },
    },

    actions: {},
    getters: {
        loggedIn(state) {
            return state.id != 0;
        },
    },

    modules: {},
});
