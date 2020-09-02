package main

import (
 //"fmt"
  "net/http"
  "encoding/json"

)


type book struct {
	Titulo string `json: titulo`
	Autor  string    `json: autor`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
  
   http.ServeFile(w, r, "ajax02.html")
}


func darMensaje(w http.ResponseWriter, r *http.Request) {

	//r.ParseForm() 
	//x := r.Form.Get("y")
  
  	//fmt.Printf("%s", x)

  

  	libro := book{
  			Titulo: "La Casa",
  			Autor:  "Paco Roca",
  		}

   
  //w.Header().Set("Content-Type", "application/json") 

     json.NewEncoder(w).Encode(libro)

}



func main() {
	
  http.HandleFunc("/dato", darMensaje)
  http.HandleFunc("/",mostrarHTML)
  
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    return
  }
}