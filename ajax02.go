package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type book struct {
	Titulo string `json: titulo`
	Autor  string `json: autor`
}
type UserData struct {
	User    string
	Mail    string
	Lastcon string
}

type SendUser struct {
	User    bool
	Mail    bool
	Lastcon bool
}

var data = [4]UserData{
	UserData{"Nick", "nicks@anymail.com", "02-28-2021 09:38:22"},
	UserData{"Luis", "luis.eds@anymail.com", "02-27-2021 19:12:02"},
	UserData{"Chris", "chris@anymail.com", "02-25-2021 12:08:12"},
	UserData{"Oscar", "the_odb@anymail.com", "02-20-2021 17:58:09"},
}

var changed = [4]SendUser{
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
}

var reset = [4]SendUser{
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
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

func cliente3(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "cliente3.html")
}

func vuejs(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "vue01.html")
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

	arreglo := [6]string{"hola", "amigo", "Como", "estas", "adelante", "pasaa"}

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		fmt.Fprintf(w, arreglo[5])
	} else if r.Method == "POST" {
		pos := r.Form.Get("y")
		dat := r.Form.Get("n")
		i, _ := strconv.Atoi(pos)

		arreglo[i%5] = dat
		fmt.Fprintf(w, "ok")

	}

}

func reto2(w http.ResponseWriter, r *http.Request) {
	arreglo := [4]string{
		"{\"Persona\":\"Nick\",\"Libro\":\"Orgullo y Prejuicio\",\"mascota\":\"manchas\"}",
		"{\"Persona\":\"Luis\",\"Libro\":\"Harry Potter\",\"mascota\":\"Sissi\"}",
		"{\"Persona\":\"Chris\",\"Libro\":\"Quiubole con\",\"mascota\":\"Nina\"}",
		"{\"Persona\":\"Oscar\",\"Libro\":\"Sapiens\",\"mascota\":\"Pirata\"}",
	}

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		pos := r.Form.Get("y")
		i, _ := strconv.Atoi(pos)
		fmt.Printf("%s\n", pos)
		fmt.Fprintf(w, arreglo[i%4])
	}
}

func reto3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		toSend := [4]UserData{} // Informacion que va a enviar realmente
		for i, v := range changed {
			var n, m, d string

			// Guardará los datos que no esten bloqueados en las variables "n, m, d"
			if !v.User {
				n = data[i].User
				v.User = true
			} else if v.User {
				n = ""
			}

			if !v.Mail {
				m = data[i].Mail
				v.Mail = true
			} else if v.Mail {
				m = ""
			}

			if !v.Lastcon {
				d = data[i].Lastcon
				v.Lastcon = true
			} else if v.Lastcon {
				d = ""
			}

			// Actualiza los permisos modificados a la variable de permisos
			changed[i] = v
			// Crea el objeto a enviar con la informacion permitida
			toSend[i] = UserData{n, m, d}

		}
		json.NewEncoder(w).Encode(toSend)
	} else if r.Method == "POST" {

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()

		id, _ := strconv.Atoi(r.Form.Get("id"))
		name := r.Form.Get("name")
		mail := r.Form.Get("mail")
		date, _ := strconv.ParseBool(r.Form.Get("date"))

		fmt.Println("Updating data...")

		// Actualiza la informacion de "data" solo cuando hay algo

		if name != "" {
			fmt.Println("name: " + data[id].User + "->" + name)
			data[id].User = name
			changed[id].User = false
		}
		if mail != "" {
			fmt.Println("mail: " + data[id].Mail + "->" + mail)
			data[id].Mail = mail
			changed[id].Mail = false
		}
		if date {
			newDate := time.Now().Format("02-01-2006 15:04:05")
			fmt.Println("date: " + data[id].Lastcon + "->" + newDate)
			data[id].Lastcon = newDate
			changed[id].Lastcon = false
		}
	}

}

func resetFunc(w http.ResponseWriter, r *http.Request) {
	changed = reset // Libera los candados al estado original
}

func main() {

	http.HandleFunc("/dato", darMensaje)
	http.HandleFunc("/", mostrarHTML)
	http.HandleFunc("/reto1", reto1)
	http.HandleFunc("/reto2", reto2)
	http.HandleFunc("/reto3", reto3)
	http.HandleFunc("/cliente", cliente)
	http.HandleFunc("/cliente2", cliente2)
	http.HandleFunc("/cliente3", cliente3)
	http.HandleFunc("/reset", resetFunc)
	http.HandleFunc("/vue", vuejs)
	fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
