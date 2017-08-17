package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"encoding/json"
)

func main() {

	m := martini.Classic()

	m.Get("/badge/:size/:color/:text1/:text2", badgeHandler)

	m.NotFound(func(res http.ResponseWriter) {
		json.NewEncoder(res).Encode(map[string]string{"message": "not found" })
		res.WriteHeader(404)
	})

	m.Run()
}

func badgeHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {

	size := params["size"]
	color := params["color"]
	text1 := params["text1"]
	text2 := params["text2"]

	banner, err := CreateBanner(text1, text2, size, color)

	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Header().Set("Content-Type", "image/svg+xml")
		generateBanner(w, banner)
	}
}
