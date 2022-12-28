// this is a sea fight multiplayer game in a console
package main

import (
	"BatailleNavale/grille"
	"fmt"
	"math/rand"
	"net/http"
)

type joueur struct {
	Pseudo string
	PORT   int
}

type adversaire struct {
	IP   string
	PORT int
}

func main() {

	fmt.Println("============= Bienvenue dans la bataille navale =============")
	var joueur joueur
	var adversaire adversaire

	fmt.Println("Entrer un pseudo visible par les autres joueurs")
	fmt.Scan(&joueur.Pseudo)
	fmt.Println("Bienvenue", joueur.Pseudo)

	fmt.Println("Entrer un port pour le serveur")
	fmt.Scan(&joueur.PORT)

	http.HandleFunc("/board", g.AfficherGrille)

	http.ListenAndServe(":"+string(rune(joueur.PORT)), nil)

	fmt.Println("Entrer l'adresse IP de l'adversaire")
	fmt.Scan(&adversaire.IP)

	fmt.Println("Entrer un port pour l'adversaire")
	fmt.Scan(&adversaire.PORT)

	//check si l'utilisateur existe
	//si oui, on recupere la grille de l'adversaire

	//recupere la grille de l'adversaire











	fmt.Print("\033[H\033[2J")

	//initialisation de la grille
	var g grille.Grille
	g.InitGrille()

	//boucle de jeu
	for !g.PartieTerminee() {

		//affiche les coordonnées des bateaux dans la grille
		g.AfficheCordBateau()

		//affichage de la grille
		g.AfficherGrille(w http.ResponseWriter, r *http.Request)

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
				g.RetirerBateau()
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


func handleBoard(w http.ResponseWriter, r *http.Request) {
	
}