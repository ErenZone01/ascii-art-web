package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./erreur/404.html")
		//http.ServeFile(w, r, "./erreur/404.html")
		template.Execute(w, r)
		return
	} else {

		template, err := template.ParseFiles("./server/index.html")
		if err != nil {
			fmt.Print("une erreur  est recu")
			return
		}

		Response := ""
		template.Execute(w, Response)
	}
	// http.ServeFile(w, r, "server")
}

type result struct {
	Response string
}


func serveur() {
	var actif = true
	


	http.HandleFunc("/", handler) // Associez la fonction de gestion à la route "/"
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	fmt.Println("Serveur en cours d'exécution sur http://localhost:8080/")
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			textes := r.FormValue("texte")
			banner := r.FormValue("radio")

			if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
				w.WriteHeader(http.StatusInternalServerError)
				template, _ := template.ParseFiles("./erreur/500.html")
				//http.ServeFile(w, r, "./erreur/500.html")
				template.Execute(w, r)
				return
			}

			textes = strings.ReplaceAll(textes, "\r\n", "\\r")
			var tabtexte = strings.Split(textes, "\\r")

			for i := 0; i < len(textes); i++ {
				for j := 0; j < len(asciiGlobal); j++ {
					if (rune(textes[i]) == asciiGlobal[j]) && rune(textes[i]) != '\r' {
						actif = true
						break
					} else {
						actif = false
					}
				}
				if !actif {
					w.WriteHeader(http.StatusBadRequest)
					template, _ := template.ParseFiles("./erreur/400.html")
					template.Execute(w, r)
					//http.ServeFile(w, r, "./erreur/400.html")
					return
				}
			}

			var ascii []string
			for i := 0; i < len(tabtexte); i++ {
				if tabtexte[i] == "\\r" {
					ascii = append(ascii, "\\r")
					continue
				}

				ascii = append(ascii, assciWeb(tabtexte[i], banner))
			}

			a := result{}

			var texte string
			for i := 0; i < len(ascii); i++ {
				texte += ascii[i]
				if i+1 != len(ascii) {
					texte += "\r\n"
				}
			}
			//fmt.Println(texte)
			a.Response = texte
			//fmt.Println(a.Response)
			//fmt.Println("Données reçues :", banner)
			//fmt.Println("Données reçues :", textes)
			template, err := template.ParseFiles("./server/index.html")
			if err != nil {
				fmt.Print("erreur reçu")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			template.Execute(w, a)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil)) // Écoutez sur le port 8080
}
