package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var errBanned = errors.New("restricted/banned")

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

	if err != nil || jerr != nil {
		app.err.Println(err)
		return "", err
	}

	if session.Priv&1 == 0 {
		return "", errBanned
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		string(js),
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(session.ID),
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(app.conf.secret))

	if err != nil {
		app.err.Println(err)
		return "", err
	}

	if err := app.DB.LastSeen(session.ID); err != nil {
		app.err.Println("Could not update last seen:", err)
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
