<template>
    <nav class="navbar" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <a
                role="button"
                class="navbar-burger"
                @click="nav = !nav"
                :class="{ 'is-active': nav }"
                aria-label="menu"
                aria-expanded="false"
                data-target="navbarmain"
            >
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
            </a>
        </div>

        <div id="navbarmain" class="navbar-menu" :class="{ 'is-active': nav }">
            <div class="navbar-start">
                <router-link class="navbar-item" to="/">Home</router-link>

                <!-- prettier-ignore -->
                <router-link class="navbar-item" to="/leaderboard?m=vn!std">Leaderboard</router-link>

                <!-- prettier-ignore -->
                <router-link class="navbar-item" to="/info">Documentation</router-link>

                <!-- prettier-ignore -->
                <router-link class="navbar-item" to="/score">Records</router-link>

                <div class="navbar-item has-dropdown is-hoverable">
                    <a class="navbar-link">More</a>

                    <div class="navbar-dropdown">
                        <a class="navbar-item" href="/info/about">About</a>
                        <a class="navbar-item" href="/info/rules">Rules</a>
                        <!-- prettier-ignore -->
                        <a class="navbar-item" href="/info/support">Support</a>

                        <hr class="navbar-divider" />
                        <a class="navbar-item" :href="source">Source code</a>
                    </div>
                </div>
            </div>

            <div class="navbar-end">
                <div class="navbar-item search-outer">
                    <input
                        type="text"
                        class="input"
                        id="search-input"
                        v-model="search"
                        placeholder="Search for a user..."
                    />

                    <ul
                        class="search-wrapper"
                        v-show="showSearch"
                        v-if="results.length"
                    >
                        <li v-for="res in results" :key="res.id">
                            <img :src="ava(res.id)" />
                            <img :src="cc(res.country)" width="20" />
                            <a :href="'/profile/' + res.id">{{
                                res.username
                            }}</a>
                        </li>
                    </ul>
                    <div
                        class="search-wrapper"
                        v-show="showSearch"
                        v-else-if="search.length > 1"
                    >
                        <LoadingMini v-if="loading" />
                        <h4 v-else>Nothing here...</h4>
                    </div>
                </div>
                <div
                    class="navbar-item has-dropdown is-hoverable"
                    v-if="$store.getters.loggedIn"
                >
                    <router-link :to="prof" class="navbar-link">
                        <img class="avatar" :src="ava($store.state.id)" />

                        {{ $store.state.username }}
                    </router-link>

                    <div class="navbar-dropdown">
                        <router-link class="navbar-item" to="/profile">
                            Settings
                        </router-link>

                        <router-link class="navbar-item" to="/profile/picture">
                            Profile Picture
                        </router-link>

                        <hr class="navbar-divider" />
                        <a class="navbar-item" @click="$store.commit('logout')">
                            Log out
                        </a>
                    </div>
                </div>
                <div class="navbar-item" v-else>
                    <div class="buttons">
                        <a href="/info/register" class="button is-primary">
                            <strong>Sign up</strong>
                        </a>
                        <router-link to="/login" class="button is-light">
                            Log in
                        </router-link>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { UserSearch } from "@/types/user";

import LoadingMini from "@/components/global/LoadingMini.vue";
import * as alert from "@/util/error";
import config from "../../../config.json";

export default defineComponent({
    components: { LoadingMini },
    data() {
        return {
            nav: false,
            loading: true,
            search: "",
            results: [] as UserSearch[],
            showSearch: false,
            timer: 0,
        };
    },

    mounted() {
        document.body.addEventListener("click", this.updateSearch);
    },

    methods: {
        async lookup(query: string) {
            this.loading = true;
            const res: UserSearch[] = await fetch(
                `${config.api}/users?s=${encodeURIComponent(query)}`
            )
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            this.loading = false;
            if (!res) return;
            this.results = res;
        },

        updateSearch(e: Event) {
            // prettier-ignore
            // eslint-disable-next-line
            this.showSearch = document.querySelector(".search-outer")!.contains(e.target as HTMLInputElement) || document.querySelector("#search-input") == document.activeElement;
        },

        ava(id: number) {
            return `${config.api}/users/${id}/avatar`;
        },

        cc(cc: string) {
            return `/img/flags/flag-${cc}.svg`;
        },
    },

    computed: {
        source(): string {
            return config.source;
        },

        prof(): string {
            return "/profile/" + this.$store.state.id;
        },
    },

    watch: {
        search(query: string) {
            clearInterval(this.timer);
            this.loading = true;

            if (query.length > 1) {
                this.timer = setTimeout(() => {
                    this.lookup(query);
                }, 500);
            } else {
                this.results = [];
            }
        },
    },
});
</script>

<style scoped>
.avatar {
    margin-right: 15px;
}

.search-wrapper {
    position: absolute;
    box-shadow: 0 1px 1px #000;
    background-color: #fff;
    width: 100%;
    top: 3rem;
    padding: 1em;
}

.search-wrapper * {
    vertical-align: middle;
    margin-left: 5px;
    overflow: hidden;
}

.search-wrapper li {
    padding: 10px 0;
}

.search-wrapper a {
    display: inline-block;
    width: 50%;
}

.search-wrapper li img:nth-child(1) {
    max-height: 40px;
}
</style>
