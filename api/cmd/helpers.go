package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/exp/slices"
)

var uptime time.Time
var countries = []string{
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
}

type simpleResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (app *application) badRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	app.JSON(w, simpleResp{
		Code:    400,
		Message: "Bad request!",
	})
}

func (app *application) internalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	app.JSON(w, simpleResp{
		Code:    500,
		Message: "Internal server error!",
	})
}

func (app *application) unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	app.JSON(w, simpleResp{
		Code:    401,
		Message: "Unauthorized!",
	})
}

func (app *application) JSON(w http.ResponseWriter, jdata any) {
	j, err := json.Marshal(jdata)

	if err != nil {
		app.err.Println("Could not parse data!")
		app.internalError(w)

		return
	}

	if string(j) == "null" {
		j = []byte("[]")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (app *application) getUptime() string {
	return time.Since(uptime).String()
}

func init() {
	uptime = time.Now()
}

func (app *application) validCountry(c string) bool {
	return slices.Contains(countries, c)
}
