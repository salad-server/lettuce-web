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
                <router-link class="navbar-item" to="/info/records">Records</router-link>

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
                <div
                    class="navbar-item has-dropdown is-hoverable"
                    v-if="$store.getters.loggedIn"
                >
                    <router-link :to="prof" class="navbar-link">
                        <img
                            class="avatar"
                            :src="`${avatar}/${$store.state.id}`"
                        />

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
import config from "../../../config.json";

export default defineComponent({
    data() {
        return {
            nav: false,
        };
    },

    computed: {
        source() {
            return config.source;
        },

        avatar() {
            return config.avatar;
        },

        prof() {
            return "/profile/" + this.$store.state.id;
        },
    },
});
</script>

<style scoped>
.avatar {
    margin-right: 15px;
}
</style>
