package main

import (
        "net/http"
				"fmt"
        )

func main() {
							//wardamos en variable fs y Definomos con http rl FileServer y le indicamos donde va a find archivos estaticos
				fs := http.FileServer(http.Dir("./static"))
				// Le indicamos al servidor http que  cuando le llegen peticiones en /static utilice  http.StripPrefix("/static/", fs))
				http.Handle("/static/", http.StripPrefix("/static/", fs))

				// Definimos la ruta que llamará la función home
				http.HandleFunc("/", home)


				// Con el segundo método hemos creado la función en el mismo lugar que creamos la ruta y usamos el paquete "fmt"(Que tenéis que agregar en la sección de los imports) para servir el contenido de una forma un poco mas elegante:
			http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) { 
 							// Paquete fmt es standar para formatear datos con go
							fmt.Fprintln(w, "Host: ",req.Host)
							// Fprintln para escribir
 							fmt.Fprintln(w, "URI: ",req.RequestURI)
 							fmt.Fprintln(w, "Method: ",req.Method)
 							fmt.Fprintln(w, "RemoteAddr: ",req.RemoteAddr)
 							})
						

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
