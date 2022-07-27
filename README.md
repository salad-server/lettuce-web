# Lettuce-web

![code-style](https://img.shields.io/badge/code_style-prettier-ff69b4.svg) ![last-commit](https://img.shields.io/github/last-commit/salad-server/lettuce-web) ![pr-closed](https://img.shields.io/github/issues-pr-closed/salad-server/lettuce-web)

The frontend for this server. Created using Vue 3. The API can be found in the `/api` directory. Internally, this project is known as **mini-web**. You can read the route information in `ROUTES.md`.

This frontend was designed to work with [bancho.py](https://github.com/osuAkatsuki/bancho.py). In theory this could work with other stacks by modifying the queries in the API, as the API follows the [standard golang project layout](https://github.com/golang-standards/project-layout).

## Differences

This frontend makes use of the Vue router to create a single page application. This means that while loading times will be faster, the frontend won't be SEO compatible.

The purpose of this project was to create a small and fast frontend that is capable of doing anything you would need from a frontend. If SEO is important to you, I don't recommend you use this project.

## Requirements
-   **Node v16.15.0 and Go 1.18** are required.
-   **make and git** are also required, but on most systems these are installed by default.
-   **MySQL and Redis** are required, but should already be installed from bancho.py.
-   **[upx](https://upx.github.io/)** is optional, but recommended as it will compress the API.

You can use [nvm](https://github.com/nvm-sh/nvm) and [gvm](https://github.com/moovweb/gvm) to install Node and Go.


## Setup

In order to setup this frontend, you must already have a functional [bancho.py](https://github.com/osuAkatsuki/bancho.py) instance running.

The first step is to clone the source code.

```sh
$ git clone https://github.com/salad-server/lettuce-web.git
$ cd lettuce-web
```

Create the configs.

```sh
$ cp config.example.json config.json
$ cp api/config.example.yaml api/config.yaml
```

Edit the configs. This should be pretty straightforward, `config.json` is mostly for visual purposes. `api/config.yaml` has some instructions written in the file itself, but if you have any issues you're welcome to open an issue and I'll get to it as soon as I can.

```sh
$ nano config.json
$ nano api/config.yaml
```

Install the dependencies and build the project.

```sh
$ npm i
$ npm run build
$ cd api
$ make build
$ make build-prod # OPTIONAL: use this instead for a smaller output. This requires upx to be installed.
$ cd ..
```

Import the SQL modifications.

```sh
$ mysql -u MY_DB_USER -p MY_DB_NAME < ext/web.sql
```

Finally, configure your nginx.

```sh
$ sudo cp ext/nginx.conf /etc/nginx/conf.d/lettuce.conf
$ sudo nano /etc/nginx/conf.d/lettuce.conf # edit this to your domain, certs, etc...
$ sudo nginx -s reload
```

Congrats! Now you should be able to access the frontend. Simply start the API and visit your server.

```sh
$ cd api
$ ./mini-api
```

You can use any process manager or something like [tmux](https://github.com/tmux/tmux) if you want to keep the API running in the background. The frontend itself doesn't need to be running as the static files are served by nginx.

**OPTIONAL:** If you want in-game redirects to work, you can include `/ext/ingame.conf` in your [bancho.py](https://github.com/osuAkatsuki/bancho.py) nginx config. Also be sure to disable `REDIRECT_OSU_URLS` for [bancho.py's](https://github.com/osuAkatsuki/bancho.py) `.env`.

**OPTIONAL:** You can install upx with:
```sh
$ sudo apt-get install upx-ucl
```

If you had any issues setting this up you're welcome to open an issue. I'll get to it as soon as I can.

## Contributing

If you're interested in contributing I'm happy to read through any pull request. You can find general information about the project in `TODO.md`, `ROUTES.md`, and `api/ROUTES.md`.

The git commit syntax I've been using is:

-   [+]/[-] - Adding/removing something
-   feat/chore - [What the commit is adding](https://gist.github.com/Cyan903/9ea365980a9beae85092a20ce2910886)
-   : - Small description

Which makes: `[+] feat: Adding something`. The syntax isn't completely required, but it would be nice if followed.

## License

  [MIT](LICENSE)
