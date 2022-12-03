package main

import (
	"fmt"
	"net/http"
	"BatailleNavale/grille"
	"math/rand"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Hello")
	case http.MethodPost:
	}
}

func main() {

	//initialisation de la grille
	var g grille.Grille
	g.InitGrille()

	//boucle de jeu
	for !g.EstTerminee() {
		//affichage de la grille
		g.AfficherGrille()

		//on demande au joueur de tirer
		var x, y int

		
		fmt.Print("Entrez les coordonnées de la case à tirer (x,y) (0 à 9) : ")
		/* fmt.Scan(&x, &y) */
		
		//simulation de saisie d'un joueur
		x = rand.Intn(10)
		y = rand.Intn(10)



		//on tire
		if g.Tirer(x, y) {
			fmt.Println("Touché !")
			if g.EstCoule(x, y) {
				fmt.Println("Coulé !")
			}
			fmt.Println("Il vous reste", g.CompterBateaux(), "bateaux à couler")

		} else {
			fmt.Println("Raté !")
		}
	}

	//affichage de la grille
	g.AfficherGrille()
	fmt.Println("Partie terminée !")

/* 
	
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":9000", nil)
 */
}