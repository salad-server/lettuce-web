<template>
    <nav class="navbar" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <a
                role="button"
                class="navbar-burger"
                aria-label="menu"
                aria-expanded="false"
                data-target="navbarmain"
            >
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
            </a>
        </div>

        <div id="navbarmain" class="navbar-menu">
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
                        <a class="navbar-item">Source code</a>
                    </div>
                </div>
            </div>

            <div class="navbar-end">
                <div class="navbar-item">
                    <div class="buttons">
                        <a class="button is-primary">
                            <strong>Sign up</strong>
                        </a>
                        <a class="button is-light"> Log in </a>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</template>
