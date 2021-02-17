package main

import (
	"fmt"
	"net/http"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
	data := [5]int{666, 420, 69, 1313, 42069}

	fmt.Fprintf(w, "data: %v", data)
	//http.ServeFile(w, r, "data")

}

func main() {

	http.HandleFunc("/", darMensaje)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	fmt.Println("Server READY!!")
	if err != nil {
		return
	}
}
