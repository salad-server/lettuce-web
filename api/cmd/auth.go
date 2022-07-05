package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func (app *application) AuthLogin(req loginRequest) (string, error) {
	bcryptPW, err := app.DB.UserCreds(req.Email)

	if err != nil {
		app.err.Println(err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(bcryptPW), []byte(req.Password)); err != nil {
		app.err.Println(err)
		return "", err
	}

	// all is good, get stats
	session, err := app.DB.UserSession(req.Email)
	js, jerr := json.Marshal(session)

	// TODO: Check if privilege is not restricted/banned.
	if err != nil || jerr != nil {
		app.err.Println(err)
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		string(js),
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(session.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(app.conf.secret))

	if err != nil {
		app.err.Println(err)
		return "", err
	}

	return token, nil
}

func (app *application) IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Token")
		token, err := jwt.ParseWithClaims(userToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.conf.secret), nil
		})

		if err != nil {
			app.err.Println(err)
			app.unauthorized(w)

			return
		}

		claims := token.Claims.(*UserClaims)
		ctx := context.WithValue(r.Context(), UserClaims{}, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
