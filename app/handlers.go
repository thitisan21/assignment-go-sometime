package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Time struct {
	Time string `json:"current_time" `
}

type Error struct {
	StatusCode int    `json:"statusCode" `
	Msg        string `json:"message" `
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {

	queryTz := r.URL.Query().Get("tz")

	if queryTz == "" {

		res := &Time{Time: time.Now().UTC().String()}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else {
		res := map[string]any{}
		for _, s := range strings.Split(queryTz, ",") {
			loc, err := time.LoadLocation(s)
			if err != nil {
				w.WriteHeader(404)
				resMsg := []byte("invalid timezone")
				w.Write(resMsg)
				return
			}

			res[s] = time.Now().In(loc).String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
