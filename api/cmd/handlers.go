package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"os"
	"strconv"

	"mini-api/internal/scores"

	"github.com/go-chi/chi/v5"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type profileRequest struct {
	Bio       string `json:"bio"`
	Country   string `json:"country"`
	Playstyle int    `json:"playstyle"`
	Email     string `json:"email"`
}

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	if app.conf.env == "development" {
		app.JSON(w, devResponse{
			Mode:   app.conf.env,
			Front:  app.conf.cors[0],
			Uptime: app.getUptime(),
		})

		return
	}

	http.Redirect(w, r, app.conf.cors[0], http.StatusPermanentRedirect)
}

func (app *application) Scores(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, ierr := strconv.Atoi(id)
	best, berr := strconv.Atoi(r.FormValue("b"))
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := (ierr != nil || berr != nil || perr != nil)
	badArgs := (uid <= -1 || best < 0 || best > 1 || page <= -1)

	if badParams || badArgs || !scores.ValidGamemode(mode) {
		app.badRequest(w)
		return
	}

	res, err := app.DB.GetScores(best != 0, uid, page, mode)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Score(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sid, err := strconv.Atoi(id)

	if err != nil {
		app.badRequest(w)
		return
	}

	res, err := app.DB.ScoreInfo(sid)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Stats(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mode := r.FormValue("m")
	uid, serr := strconv.Atoi(id)

	if serr != nil || !scores.ValidGamemode(mode) {
		app.badRequest(w)
		return
	}

	res, err := app.DB.UserStats(uid, mode)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Info(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, serr := strconv.Atoi(id)

	if serr != nil {
		app.badRequest(w)
		return
	}

	res, err := app.DB.UserInfo(uid)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Beatmap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, serr := strconv.Atoi(id)

	if serr != nil {
		app.badRequest(w)
		return
	}

	res, err := app.DB.BeatmapInfo(bid)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) BeatmapSets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, serr := strconv.Atoi(id)

	if serr != nil {
		app.badRequest(w)
		return
	}

	res, err := app.DB.BeatmapSets(bid)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Records(w http.ResponseWriter, r *http.Request) {
	res, err := app.DB.Records()

	if err != nil {
		app.err.Println(err)
		app.internalError(w)
		return
	}

	app.JSON(w, res)
}

func (app *application) BeatmapLeaderboard(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, berr := strconv.Atoi(id)
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := berr != nil || perr != nil || !scores.ValidGamemode(mode)
	badArgs := (bid <= -1 || page <= -1)

	if badParams || badArgs {
		app.badRequest(w)
		return
	}

	res, err := app.DB.BeatmapLeaderboard(bid, page, mode)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Leaderboard(w http.ResponseWriter, r *http.Request) {
	page, perr := strconv.Atoi(r.FormValue("p"))
	sort, serr := strconv.Atoi(r.FormValue("s"))
	mode := r.FormValue("m")

	if perr != nil || serr != nil || page <= -1 || !scores.ValidGamemode(mode) {
		app.badRequest(w)
		return
	}

	// 0 = pp
	// 1 = score
	res, err := app.DB.Leaderboard(sort != 1, mode, page)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) DocsListing(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(app.conf.docs)
	listing := make([]string, len(files))

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	for i, f := range files {
		listing[i] = f.Name()
	}

	app.JSON(w, listing)
}

func (app *application) Docs() http.Handler {
	return http.StripPrefix(
		"/docs/",
		http.FileServer(http.Dir(app.conf.docs)),
	)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var res loginRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&res); err != nil {
		app.badRequest(w)
		return
	}

	if len(res.Email) < 3 || len(res.Password) < 3 || len(res.Email) > 100 || len(res.Password) > 100 {
		app.badRequest(w)
		return
	}

	res.Password = fmt.Sprintf("%x", md5.Sum([]byte(res.Password)))
	token, err := app.AuthLogin(res)

	if err != nil {
		if errors.Is(err, errBanned) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)

			app.JSON(w, simpleResp{
				Code:    403,
				Message: "You are restricted/banned!",
			})

			return
		}

		app.unauthorized(w)
		return
	}

	w.Header().Set("Content-Type", "text/raw")
	w.Write([]byte(token))
}

func (app *application) WhoAmI(w http.ResponseWriter, r *http.Request) {
	if app.conf.env == "development" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(r.Context().Value(UserClaims{}).(*UserClaims).User))
	} else {
		app.unauthorized(w)
	}
}

func (app *application) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	uid, _ := strconv.Atoi(id)
	profile, err := app.DB.ProfileEdit(uid)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, profile)
}

func (app *application) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var res profileRequest

	id := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	uid, _ := strconv.Atoi(id)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&res); err != nil {
		app.badRequest(w)
		return
	}

	_, merr := mail.ParseAddress(res.Email)

	// bio cannot be 2,000 characters
	// country must be valid
	// playstyle... is the playstyle
	// email must be valid
	if len(res.Bio) > 2000 || !app.validCountry(res.Country) || merr != nil {
		app.badRequest(w)
		return
	}

	if app.DB.EmailExists(uid, res.Email) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)

		app.JSON(w, simpleResp{
			Code:    409,
			Message: "email is already in use!",
		})

		return
	}

	if err := app.DB.ProfileUpdate(res.Bio, res.Country, res.Email, res.Playstyle, uid); err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.success(w)
}

func (app *application) ProfilePicture(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(app.conf.profile.max)

	uid := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	file, handler, err := r.FormFile("file")

	if err != nil || handler.Size > app.conf.profile.max {
		app.badRequest(w)
		return
	}

	buff, err := ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	if ftype := http.DetectContentType(buff); ftype != "image/png" && ftype != "image/jpeg" {
		app.err.Println(handler.Filename, "is not a png/jpeg!")
		app.badRequest(w)

		return
	}

	t, err := os.Create(fmt.Sprintf("%s/%s.jpg", app.conf.profile.path, uid))
	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	defer t.Close()

	t.Write(buff)
	app.success(w)
}

func (app *application) ProfilePictureRemove(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	img := fmt.Sprintf("%s/%s.jpg", app.conf.profile.path, uid)

	if _, err := os.Stat(img); errors.Is(err, os.ErrNotExist) {
		app.badRequest(w)
		return
	}

	if err := os.Remove(img); err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.success(w)
}

func (app *application) Pinned(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, ierr := strconv.Atoi(id)
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := (ierr != nil || perr != nil)
	badArgs := (uid <= -1 || page <= -1)

	if badParams || badArgs || !scores.ValidGamemode(mode) {
		app.badRequest(w)
		return
	}

	res, err := app.DB.GetPinned(uid, page, mode)

	if err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Pin(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	sid, err := strconv.Atoi(r.FormValue("s"))
	uid, _ := strconv.Atoi(id)

	if err != nil {
		app.badRequest(w)
		return
	}

	if !app.DB.ScoreExists(uid, sid) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		app.JSON(w, simpleResp{
			Code:    404,
			Message: "score does not exist!",
		})

		return
	}

	if app.DB.PinnedExists(sid) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)

		app.JSON(w, simpleResp{
			Code:    409,
			Message: "score is already pinned!",
		})

		return
	}

	if err := app.DB.PinScore(sid); err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.success(w)
}

func (app *application) Unpin(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(UserClaims{}).(*UserClaims).Issuer
	sid, err := strconv.Atoi(r.FormValue("s"))
	uid, _ := strconv.Atoi(id)

	if err != nil {
		app.badRequest(w)
		return
	}

	// ensure we own the score and we aren't restricted
	if !app.DB.ScoreExists(uid, sid) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		app.JSON(w, simpleResp{
			Code:    404,
			Message: "score does not exist!",
		})

		return
	}

	if !app.DB.PinnedExists(sid) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)

		app.JSON(w, simpleResp{
			Code:    409,
			Message: "score isn't pinned!",
		})

		return
	}

	if err := app.DB.UnpinScore(sid); err != nil {
		app.err.Println(err)
		app.internalError(w)

		return
	}

	app.success(w)
}
