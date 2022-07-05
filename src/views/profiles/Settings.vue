<template>
    <Loading v-if="loading" />
    <div v-else>
        <div class="columns is-centered">
            <div class="my-8 column dcard is-6">
                <form @submit.prevent="update()" class="mx-5">
                    <!-- prettier-ignore -->
                    <h1 class="has-text-centered my-5 is-size-3 has-text-weight-bold">Profile Settings</h1>

                    <div class="field">
                        <label class="label">Email</label>
                        <div class="control">
                            <input
                                :class="validEmail"
                                type="email"
                                v-model="email"
                                placeholder="Enter a new email"
                            />
                        </div>
                    </div>

                    <div class="field">
                        <label class="label"> Country </label>
                        <div class="control flag-control">
                            <div class="select">
                                <select v-model="country">
                                    <option
                                        v-for="c in flags"
                                        :key="c"
                                        :value="c"
                                    >
                                        {{ c.toUpperCase() }}
                                    </option>
                                </select>
                            </div>
                            <img :src="flag" width="50" />
                        </div>
                    </div>

                    <div class="field">
                        <div class="control">
                            <textarea
                                class="textarea"
                                v-model="bio"
                                maxlength="2000"
                                placeholder="About yourself..."
                            ></textarea>
                            <label class="label count">
                                {{ bio.length }}/2000
                            </label>
                        </div>
                    </div>

                    <div class="field is-grouped columns is-centered my-5">
                        <button
                            class="button is-success"
                            @click="update"
                            :disabled="!valid"
                        >
                            Update
                        </button>
                    </div>
                </form>
            </div>
        </div>

        <!-- TODO: Playstyle -->
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ProfileSettings } from "@/types/user";

import Loading from "@/components/global/Loading.vue";
import Swal from "sweetalert2";
import * as alert from "@/util/error";
import config from "../../../config.json";

// https://stackoverflow.com/a/46181
const val = (email: string) => {
    return email
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
};

// prettier-ignore
const countries = [
	"ad", "ae",
	"af", "ag", "ai", "al", "am", "an",
	"ao", "aq", "ar", "as", "at", "au",
	"aw", "ax", "az", "ba", "bb", "bd",
	"be", "bf", "bg", "bh", "bi", "bj",
	"bl", "bm", "bn", "bo", "bq", "br",
	"bs", "bt", "bv", "bw", "by", "bz",
	"ca", "cc", "cd", "cf", "cg", "ch",
	"ci", "ck", "cl", "cm", "cn", "co",
	"cr", "cs", "cu", "cv", "cw", "cx",
	"cy", "cz", "de", "dj", "dk", "dm",
	"do", "dz", "ec", "ee", "eg", "eh",
	"er", "es", "et", "eu", "fi", "fj",
	"fk", "fm", "fo", "fr", "ga", "gb",
	"gd", "ge", "gf", "gg", "gh", "gi",
	"gl", "gm", "gn", "gp", "gq", "gr",
	"gs", "gt", "gu", "gw", "gy", "hk",
	"hm", "hn", "hr", "ht", "hu", "id",
	"ie", "il", "im", "in", "io", "iq",
	"ir", "is", "it", "je", "jm", "jo",
	"jp", "ke", "kg", "kh", "ki", "km",
	"kn", "kp", "kr", "kw", "ky", "kz",
	"la", "lb", "lc", "li", "lk", "lr",
	"ls", "lt", "lu", "lv", "ly", "ma",
	"mc", "md", "me", "mf", "mg", "mh",
	"mk", "ml", "mm", "mn", "mo", "mp",
	"mq", "mr", "ms", "mt", "mu", "mv",
	"mw", "mx", "my", "mz", "na", "nc",
	"ne", "nf", "ng", "ni", "nl", "no",
	"np", "nr", "nu", "nz", "om", "pa",
	"pe", "pf", "pg", "ph", "pk", "pl",
	"pm", "pn", "pr", "ps", "pt", "pw",
	"py", "qa", "re", "ro", "rs", "ru",
	"rw", "sa", "sb", "sc", "sd", "se",
	"sg", "sh", "si", "sj", "sk", "sl",
	"sm", "sn", "so", "sr", "ss", "st",
	"sv", "sx", "sy", "sz", "tc", "td",
	"tf", "tg", "th", "tj", "tk", "tl",
	"tm", "tn", "to", "tr", "tt", "tv",
	"tw", "tz", "ua", "ug", "um", "us",
	"uy", "uz", "va", "vc", "ve", "vg",
	"vi", "vn", "vu", "wf", "ws", "xk",
	"ye", "yt", "za", "zm", "zw", "xx",
];

export default defineComponent({
    components: { Loading },
    data() {
        return {
            bio: "",
            country: "ca",
            playstyle: 0,
            email: "",

            valid: true,
            loading: true,
        };
    },

    async mounted() {
        if (!this.$store.getters.loggedIn) {
            this.$router.push("/login");
            return;
        }

        const res: ProfileSettings = await fetch(`${config.api}/@me/profile`, {
            headers: {
                Token: this.$store.state.token,
            },
        })
            .then((j) => j.json())
            .catch(() => {
                alert.API();
            });

        if (!res) return;

        this.loading = false;
        this.bio = res.bio;
        this.country = res.country;
        this.playstyle = res.playstyle;
        this.email = res.email;
    },

    methods: {
        async update() {
            const res: ProfileSettings = await fetch(
                `${config.api}/@me/profile`,
                {
                    method: "POST",
                    headers: {
                        Token: this.$store.state.token,
                    },

                    body: JSON.stringify({
                        bio: this.bio,
                        country: this.country,
                        playstyle: this.playstyle,
                        email: this.email,
                    }),
                }
            )
                .then((j) => j.json())
                .catch(() => {
                    alert.API();
                });

            if (!res) return;
            if (res?.code != 200) {
                Swal.fire({
                    title: "Invalid request!",
                    text: "Double check your inputs?",
                    icon: "error",
                });

                return;
            }

            Swal.fire({
                title: "Success!",
                icon: "success",
            });

            this.$router.push(`/profile/${this.$store.state.id}?g=vn&m=std`);
        },
    },

    computed: {
        flags() {
            return countries;
        },

        flag(): string {
            return `/img/flags/flag-${this.country}.svg`;
        },

        validEmail(): string {
            return this.valid ? "input" : "input is-danger";
        },
    },

    watch: {
        email(v: string) {
            this.valid = !!val(v);
        },
    },
});
</script>

<style scoped>
.flag-control * {
    display: inline-block;
    vertical-align: middle;
}

.flag-control img {
    margin-left: 10px;
}

.count {
    float: right;
    color: #333;
}

textarea {
    resize: none;
}
</style>
