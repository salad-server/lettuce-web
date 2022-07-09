# Lettuce-web/api

![code-style](https://img.shields.io/badge/code_style-prettier-ff69b4.svg) ![last-commit](https://img.shields.io/github/last-commit/salad-server/lettuce-web) ![pr-closed](https://img.shields.io/github/issues-pr-closed/salad-server/lettuce-web)

The API for the frontend. Follows the [standard golang project layout](https://github.com/golang-standards/project-layout). Route information can be found in `ROUTES.md` or `cmd/routes.go`. Uses [go-chi](https://github.com/go-chi/chi) and [jwt](https://github.com/golang-jwt/jwt) for routing/authentication.

Using **Go 1.18**

## Makefile
To build:
```sh
$ make build
$ make build-prod # to compress with upx
```

To run without building:
```sh
$ make serve
```

To kill:
```sh
$ make kill
```
