package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RedirectList struct {
	Redirects []Redirect `json:"list"`
}

type RedirectYAML struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

var config string

func main() {

	flag.StringVar(&config, "type", "json", "Type of config file")
	flag.Parse()

	http.HandleFunc("/", handle)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Failed to build server")
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	redirects := parserMap(config)
	http.Redirect(w, r, redirects[strings.TrimSpace(r.URL.Path[1:])], 301)
}

func parserMap(config string) map[string]string {

	switch config {
	case "json":
		fmt.Println("json")
		jsonMapFile, err := ioutil.ReadFile("../config/map.json")
		if err != nil {
			fmt.Printf("Failed to open %v \n", err)
		}
		return JSONtoRedirect(jsonMapFile)
	case "yaml":
		fmt.Println("yaml")
		jsonMapFile, err := ioutil.ReadFile("../config/map.yaml")
		if err != nil {
			fmt.Printf("Failed to open %v \n", err)
		}
		return YAMLtoRedirect(jsonMapFile)
	default:
		fmt.Println("nada")
		return make(map[string]string)
	}
}

func parserYAMLMap() map[string]string {
	jsonMapFile, err := ioutil.ReadFile("../config/map.yaml")
	if err != nil {
		fmt.Printf("Failed to open %v \n", err)
	}

	return JSONtoRedirect(jsonMapFile)

}

func YAMLtoRedirect(jsonData []byte) map[string]string {

	var redirects []RedirectYAML
	yaml.Unmarshal(jsonData, &redirects)

	mapa := make(map[string]string)

	for _, v := range redirects {
		mapa[v.From] = v.To
	}
	return mapa
}

func JSONtoRedirect(jsonData []byte) map[string]string {

	var redirects RedirectList
	yaml.Unmarshal(jsonData, &redirects)

	mapa := make(map[string]string)

	for _, v := range redirects.Redirects {
		mapa[v.From] = v.To
	}
	return mapa
}
