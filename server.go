package main

import (
        "net/http"
        )

func main() {
							//wardamos en variable fs y Definomos con http rl FileServer y le indicamos donde va a find archivos estaticos
				fs := http.FileServer(http.Dir("./static"))
				// Le indicamos al servidor http que  cuando le llegen peticiones en /static utilice  http.StripPrefix("/static/", fs))
				http.Handle("/static/", http.StripPrefix("/static/", fs))

        http.ListenAndServe(":8080", nil)
        }