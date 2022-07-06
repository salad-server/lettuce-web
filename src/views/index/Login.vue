<template>
    <div class="columns is-centered">
        <div class="my-8 column dcard is-5">
            <form @submit.prevent="submit()" class="mx-5">
                <!-- prettier-ignore -->
                <h1 class="has-text-centered my-5 is-size-2 has-text-weight-bold">Log in</h1>

                <div class="field">
                    <label class="label">Email</label>
                    <div class="control">
                        <input
                            :class="valid(invalid.email)"
                            type="email"
                            v-model="email"
                            placeholder="Enter your email"
                        />
                    </div>
                </div>

                <div class="field">
                    <label class="label">Password</label>
                    <div class="control">
                        <input
                            :class="valid(invalid.password)"
                            type="password"
                            v-model="password"
                            placeholder="Enter your password"
                        />
                    </div>
                </div>

                <div class="field is-grouped columns is-centered my-5">
                    <div class="control">
                        <button class="button is-success" :disabled="validForm">
                            Log in
                        </button>
                    </div>
                    <div class="control">
                        <router-link
                            to="/info/forgot"
                            class="button is-link is-light"
                        >
                            I forgot my password!
                        </router-link>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import * as alert from "@/util/error";
import Swal from "sweetalert2";
import config from "../../../config.json";

// https://stackoverflow.com/a/46181
const val = (email: string) => {
    return email
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
};

export default defineComponent({
    data() {
        return {
            email: "",
            password: "",

            invalid: {
                email: true,
                password: true,
                loading: false,
            },
        };
    },

    methods: {
        async submit() {
            this.invalid.loading = true;

            // prettier-ignore
            const res = await fetch(`${config.api}/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },

                body: JSON.stringify({
                    email: this.email,
                    password: this.password,
                }),
            }).catch(() => alert.API());

            this.invalid.loading = false;

            if (!res) return;
            if (res.status != 200) {
                Swal.fire({
                    title: "Unauthorized!",
                    text: "Double check your inputs?",
                    icon: "question",
                });

                return;
            }

            const token = await res.text();

            this.$store.commit("login", token);
            this.$router.push("/");
        },

        valid(n: boolean) {
            return n ? "input is-danger" : "input";
        },
    },

    watch: {
        email(mail: string) {
            if (val(mail)) {
                this.invalid.email = false;
                return;
            }

            this.invalid.email = true;
        },

        password(pw: string) {
            this.invalid.password = pw.length < 3;
        },
    },

    computed: {
        validForm(): boolean {
            return (
                this.invalid.password ||
                this.invalid.email ||
                this.invalid.loading
            );
        },
    },
});
</script>
