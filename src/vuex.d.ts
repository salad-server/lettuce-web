import { Store } from "vuex";

declare module "@vue/runtime-core" {
    interface State {
        token: string;
        username: string;
        expire: number;
        priv: number;
        id: number;
    }

    interface ComponentCustomProperties {
        $store: Store<State>;
    }
}
