package main

import (
	"fmt"
	"os"
	"strings"
)

var asciiGlobal []rune

func assciWeb(arg1 string, arg2 string) string {
	file, err := os.ReadFile(arg2 + ".txt")
	if err != nil {
		fmt.Println("le fichier est introuvable")
		os.Exit(0)
	}

	texte := string(file)

	var arguments = strings.Split(arg1, "\\r")

	var result string

	for i := 0; i < len(arguments); i++ {
		if len(arguments[i]) == 0 {
			fmt.Println()
			continue
		}
		var tab []int
		var tabPhrase []string
		var phrase string = ""
		var tabASCII []rune
		var position []int
		var textePhrase []string
		var retourLine []int
		if arguments[i] != "\r" {
			arg1 = arguments[i]
			for i := 0; i < 95; i++ {
				tab = append(tab, i)
			}
			for i := ' '; i <= '~'; i++ {
				tabASCII = append(tabASCII, i)
			}
			asciiGlobal = tabASCII
			var actif bool = false
			var compteurArg []int
			var special bool = true

			for i := 0; i < len(arg1); i++ {
				for j := 0; j < len(tabASCII); j++ {
					if i < len(arg1)-1 && (string(arg1[i]) == "\\" && string(arg1[i+1]) == "r") {
						compteurArg = append(compteurArg, i)
						retourLine = append(retourLine, i)
						retourLine = append(retourLine, i+1)
						actif = true
						break
					} else {
						if actif == false {
							if rune(arg1[i]) == tabASCII[j] {
								compteurArg = append(compteurArg, i)
								position = append(position, tab[j])
								special = true
								break
							}
						} else {
							actif = false
							break
						}

					}

				}
				if !special {
					fmt.Println("Ce programme ne prend que les carcateres de la table ASCII")
					os.Exit(0)
				}
			}

			var compteur int = 0

			for i := 0; i < len(texte); i++ {
				if texte[i] != '\n' {
					phrase += string(texte[i])
				} else {
					compteur++
					if compteur == 9 {
						tabPhrase = append(tabPhrase, phrase)
						phrase = ""
						compteur = 0
						continue
					}
					if compteur != 9 {
						phrase += "\n"
					}

				}

			}

			for i := 0; i < len(position); i++ {
				for j := 0; j < len(tabPhrase); j++ {
					if position[i] == j {
						textePhrase = append(textePhrase, tabPhrase[j])
					}
				}
			}
			x := [9][]string{}
			for _, text := range textePhrase {
				for i, y := range strings.Split(text, "\n") {
					x[i] = append(x[i], y)
				}
			}
			var texte2 string

			for i, ligne := range x {
				if i > 0 {
					for _, part := range ligne {
						if len(part) > 0 {
							part = part[:len(part)-1]
						}
						texte2 += part
						texte2 += " "
					}

					texte2 += "\n"
				}
			}
			result = texte2
		}
	}
	return result
}
