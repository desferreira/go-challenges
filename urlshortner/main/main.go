package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RedirectList struct {
	Redirects []Redirect `json:"list"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	redirects := parserJSONMap()
	http.Redirect(w, r, redirects[strings.TrimSpace(r.URL.Path[1:])], 301)
}

func parserJSONMap() map[string]string {
	jsonMapFile, err := ioutil.ReadFile("config/map.json")
	if err != nil {
		fmt.Printf("Failed to open %v \n", err)
	}

	return JSONtoRedirect(jsonMapFile)

}

func JSONtoRedirect(jsonData []byte) map[string]string {

	var redirects RedirectList
	json.Unmarshal(jsonData, &redirects)

	mapa := make(map[string]string)

	for _, v := range redirects.Redirects {
		mapa[v.From] = v.To
	}
	return mapa
}
