package main

import (
	"fmt"

	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/intro"
	"Projet-Red_GoonLegend/menu"
)

// fonction principale
func main() {
	// intro
	intro.RunIntro()

	// demander le nom
	var nom string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&nom)

	// 3. Crée le personnage
	hero := character.InitCharacter(
		nom, // nom en question
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion", "Potion", "Potion"},
	)

	// lancement du menu (appel du package Menu)
	menu.Run(&hero)
}
