<template>
    <div class="hero-body">
        <div class="container has-text-centered">
            <div class="columns is-vcentered">
                <div class="column is-5">
                    <figure class="image is-4by2">
                        <img :src="landing" alt="Landing Cover" />
                    </figure>
                </div>
                <div class="column is-5 is-offset-1">
                    <h1 class="title is-2">Welcome to {{ serverName }}</h1>
                    <h2 class="subtitle is-4">osu! private server</h2>
                    <br />
                    <p class="has-text-centered">
                        <router-link
                            to="/info/connect"
                            class="button is-medium is-info is-outlined mr-3"
                        >
                            Connection Guide
                        </router-link>

                        <router-link
                            to="/info/register"
                            class="button is-medium is-info is-outlined"
                        >
                            Register
                        </router-link>
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import config from "../../config.json";

export default defineComponent({
    computed: {
        landing(): string {
            return config.landing;
        },

        serverName(): string {
            return config.name;
        },
    },
});
</script>
