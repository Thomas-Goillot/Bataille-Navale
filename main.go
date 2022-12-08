package main

import (
	"BatailleNavale/grille"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("\033[H\033[2J")

	//initialisation de la grille
	var g grille.Grille
	g.InitGrille()

	//boucle de jeu
	for !g.PartieTerminee() {
		//affichage de la grille
		g.AfficherGrille()

		//on demande au joueur de tirer
		var x, y int

		//simulation de saisie d'un joueur
		x = rand.Intn(10)
		y = rand.Intn(10)

		fmt.Print("Entrez les coordonnées de la case à tirer (x,y) (0 à 9) : ")
		fmt.Scan(&x, &y)

		//on tire
		if g.Tirer(x, y) {
			fmt.Println("Touché !")
			if g.EstCoule(x, y) {
				fmt.Println("Coulé !")
			}
			fmt.Println("Il vous reste", g.NbBateauxRestants(), "bateaux à couler")

		} else {
			fmt.Println("Raté !")
		}
	}

	//affichage de la grille
	g.AfficherGrille()
	fmt.Println("Partie terminée !")
}
