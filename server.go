package main

import (
        "net/http"
        )

func main() {
							//wardamos en variable fs y Definomos con http rl FileServer y le indicamos donde va a find archivos estaticos
				fs := http.FileServer(http.Dir("./static"))
				// Le indicamos al servidor http que  cuando le llegen peticiones en /static utilice  http.StripPrefix("/static/", fs))
				http.Handle("/static/", http.StripPrefix("/static/", fs))

				// Definimos la ruta que llamará la función home
				http.HandleFunc("/", home)

        http.ListenAndServe(":8080", nil)
        }

/* Definimos la función home ||  objeto request que esta en la variable r y  tiene toda la informacion de la peticion que nos han hecho*/
func home(w http.ResponseWriter, r *http.Request) {
        html := "<html>";
        html += "<body>";
        html += "<h1>Hola Mundo</h1>";
        html += "</body>";
        html += "</html>";
				// Convertimos el string a un array de bytes
        w.Write( []byte(html) )
        }
