import { createStore } from "vuex";
import { AuthInfo, AuthUser } from "@/types/user";
import Swal from "sweetalert2";

let timer = 0;

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

            // every minute, check if user's session has expired
            timer = setInterval(() => {
                if (Math.floor(Date.now() / 1000) > state.expire) {
                    Swal.fire({
                        title: "Unauthorized!",
                        text: "Double check your inputs?",
                        icon: "question",
                    });

                    setTimeout(() => {
                        // eslint-disable-next-line
                        (this as any).commit("logout");
                    }, 2000);

                    clearInterval(timer);
                }
            }, 60000);
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
