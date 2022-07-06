<template>
    <div class="columns is-centered">
        <div class="my-8 column dcard is-6">
            <form @submit.prevent="" class="mx-5">
                <!-- prettier-ignore -->
                <h1 class="has-text-centered my-5 is-size-3 has-text-weight-bold">Profile Picture</h1>

                <div class="field has-text-centered">
                    <img :src="preview || pfp" alt="Profile Picture" />
                </div>

                <div class="field is-grouped columns is-centered my-5">
                    <div class="control">
                        <div class="file has-name">
                            <label class="file-label">
                                <input
                                    class="file-input"
                                    type="file"
                                    name="file"
                                    @change="updatePreview"
                                />
                                <span class="file-cta">
                                    <span class="file-icon">
                                        <img
                                            src="/img/icons/upload.svg"
                                            alt="Upload icon"
                                        />
                                    </span>
                                    <span class="file-label">
                                        Choose a file…
                                    </span>
                                </span>
                                <span class="file-name">{{ name }}</span>
                            </label>
                        </div>

                        <h3 class="has-text-centered mt-4">
                            File must be either PNG or JPG
                        </h3>
                    </div>
                </div>

                <div class="field is-grouped columns is-centered my-5">
                    <button
                        class="button is-success mx-2"
                        @click="upload"
                        :disabled="!preview"
                    >
                        Upload
                    </button>

                    <button class="button is-danger mx-2" @click="remove">
                        Remove Current
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { SimpleResponse } from "@/types/general";

import Swal from "sweetalert2";
import * as alert from "@/util/error";
import config from "../../../config.json";

export default defineComponent({
    data() {
        return {
            preview: "",
            name: "Upload a file...",
        };
    },

    async mounted() {
        if (!this.$store.getters.loggedIn) {
            this.$router.push("/login");
            return;
        }
    },

    methods: {
        async upload() {
            // prettier-ignore
            const input = document.querySelector("input[type='file']") as HTMLInputElement;
            const data = new FormData();

            if (!input.files) {
                return;
            }

            data.append("file", input.files[0]);
            data.append("user", "hubot");

            const res: SimpleResponse = await fetch(`${config.api}/@me/pfp`, {
                method: "POST",
                headers: {
                    Token: this.$store.state.token,
                },

                body: data,
            })
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            if (res.code != 200) {
                Swal.fire({
                    title: "Invalid request!",
                    text: "Did you upload a proper image?",
                    icon: "error",
                });

                return;
            }

            Swal.fire({
                title: "Success!",
                text: "You may need to reload for the changes to show...",
                icon: "success",
            });

            this.$router.push(`/profile/${this.$store.state.id}?g=vn&m=std`);
        },

        updatePreview(e: Event) {
            const input = e.target as HTMLInputElement;
            console.log(2);

            if (input.files) {
                this.preview = URL.createObjectURL(input.files[0]);
                this.name = input.files[0].name;
            }
        },

        async remove() {
            const res: SimpleResponse = await fetch(`${config.api}/@me/pfp`, {
                method: "DELETE",
                headers: {
                    Token: this.$store.state.token,
                },
            })
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            if (res.code != 200) {
                Swal.fire({
                    title: "You are using the default picture!",
                    text: "You can only remove a custom photo.",
                    icon: "error",
                });

                return;
            }

            Swal.fire({
                title: "Success!",
                text: "You may need to reload for the changes to show...",
                icon: "success",
            });

            this.$router.push(`/profile/${this.$store.state.id}?g=vn&m=std`);
        },
    },

    computed: {
        pfp() {
            return `${config.avatar}/${this.$store.state.id}`;
        },
    },
});
</script>
