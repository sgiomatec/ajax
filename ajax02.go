package main

import (
 "fmt"
  "net/http"
  "encoding/json"
  "strconv"

)


type book struct {
	Titulo string `json: titulo`
	Autor  string    `json: autor`
}
type data struct {
	data string `json: data`

}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
  
   http.ServeFile(w, r, "ajax02.html")
}
func cliente(w http.ResponseWriter, r *http.Request) {
  
  http.ServeFile(w, r, "cliente1.html")
}
func cliente2(w http.ResponseWriter, r *http.Request) {
  
  http.ServeFile(w, r, "cliente2.html")
}

func darMensaje(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() 
	x := r.Form.Get("y")
  
  	fmt.Printf("%s", x)

  

  	libro := book{
  			Titulo: "La Casa",
  			Autor:  "Paco Roca",
  		}

   
  //w.Header().Set("Content-Type", "application/json") 

     json.NewEncoder(w).Encode(libro)

}
func reto1(w http.ResponseWriter, r *http.Request) {
 
  arreglo := [6]string{"hola","amigo","Como","estas","adelante", "pasaa"}
  

	r.ParseForm() 
  w.Header().Set("Content-Type", "application/json") 
  if r.Method == "GET"{
    fmt.Fprintf(w,arreglo[5])
  } else if r.Method == "POST"{
    pos := r.Form.Get("y")
    dat := r.Form.Get("n")
    i , _:=strconv.Atoi(pos)

    arreglo[i%5]=dat
    fmt.Fprintf(w,"ok")

  }
	
}
func reto2(w http.ResponseWriter, r *http.Request) {
 
  arreglo := [4]string{"{\"Persona\":\"Nick\",\"Libro\":\"Orgullo y Prejuicio\",\"mascota\":\"manchas\"}",
                       "{\"Persona\":\"Luis\",\"Libro\":\"Harry Potter\",\"mascota\":\"Sissi\"}",
                       "{\"Persona\":\"Chris\",\"Libro\":\"Quiubole con\",\"mascota\":\"Nina\"}",
                       "{\"Persona\":\"Oscar\",\"Libro\":\"Sapiens\",\"mascota\":\"Pirata\"}"}
  

	r.ParseForm() 
  w.Header().Set("Content-Type", "application/json") 
  
    /* index := r.Form.Get("y")
    fmt.Printf("Position: %s\n",index);
    i , _:=strconv.Atoi(index)
    fmt.Printf("value: %d\n",i);

    fmt.Fprintf(w,arreglo[i%4]) */
    if r.Method == "POST"{

      pos := r.Form.Get("y")
    
      i , _:=strconv.Atoi(pos)
      fmt.Printf("%s\n",pos)
      fmt.Fprintf(w,arreglo[i%4])
  
    }
  
	
}



func main() {
	
  http.HandleFunc("/dato", darMensaje)
  http.HandleFunc("/",mostrarHTML)
  http.HandleFunc("/reto1",reto1)
  http.HandleFunc("/reto2",reto2)
  http.HandleFunc("/cliente",cliente)
  http.HandleFunc("/cliente2",cliente2)
  
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    return
  }
}