package main

import (
  "fmt"
  
  "net/http"

)



func darMensaje(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%s\n", r.Referer())
  //fmt.Printf("%s\n", r.URL)
  //fmt.Printf("%s", r.Method)
  //fmt.Printf("%s", r.UserAgent())
  //fmt.Fprintf(w, "something at the array: %d! \n'", data[4])
  //http.ServeFile(w, r, "ajax02.html")



}

func main() {
  http.HandleFunc("/", darMensaje)
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    return
  }
}