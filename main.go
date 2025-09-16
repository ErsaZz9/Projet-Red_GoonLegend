package main

import (
	"Projet-Red_GoonLegend/Menu/intro"
	"Projet-Red_GoonLegend/character"
	"fmt"
)

func main() {
	intro.RunIntro()

	// Input pour le nom
	var nom string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&nom)

	// 3. Créer le personnage
	hero := character.InitCharacter(
		nom,
		"Elfe",                                 // Classe
		1,                                      // Niveau
		100,                                    // PV max
		40,                                     // PV actuels
		[]string{"Potion", "Potion", "Potion"}, // Inventaire
	)

	// 4. Afficher le perso
	fmt.Println("Personnage créé :", hero)
}
