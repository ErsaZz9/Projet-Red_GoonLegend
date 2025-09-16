package main

import (
	"fmt"

	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/intro"
)

func main() {
	// 1. Affiche l’intro
	intro.RunIntro()

	// 2. Input pour le nom du perso
	var nom string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&nom)

	// 3. Crée le personnage
	hero := character.InitCharacter(
		nom,                                    // Nom (saisi par l’utilisateur)
		"Elfe",                                 // Classe
		1,                                      // Niveau
		100,                                    // PV max
		40,                                     // PV actuels
		[]string{"Potion", "Potion", "Potion"}, // Inventaire
	)
	// 4. Affiche les infos du perso
	hero.DisplayInfo()
}
