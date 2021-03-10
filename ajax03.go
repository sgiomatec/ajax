package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func vuejs(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "vue01.html")
}

type Film struct {
	Title    string `json: title`
	Director string `json: director`
	Year     int    `json: year`
}

func sendFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var films = [4]Film{
		Film{"No Manches Frida 2", "Nacho G. Velilla", 2016},
		Film{"Star Wars 6", "George Lucas", 1976},
		Film{"Harry Potter 3", "Alfonso Cuarón", 2002},
		Film{"Elba lazo", "Gillermo del Toro", 2009},
	}

	json.NewEncoder(w).Encode(films)
}

func main() {
	http.HandleFunc("/vue", vuejs)
	http.HandleFunc("/films", sendFilms)
	fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
