package main

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/intro"
	"Projet-Red_GoonLegend/menu"
	"Projet-Red_GoonLegend/save"
	"fmt"
	"time"
)

func main() {
	// intro
	intro.RunIntro()

	// on essaye d'abord de check si il y a une sauvegarde
	hero, err := save.LoadCharacter("save.json")
	if err != nil {
		fmt.Println("Aucune sauvegarde existante")
		time.Sleep(2 * time.Second)
		fmt.Println("Création d'un nouveau héros")

		// ⚡ utilise directement la fonction de création
		hero = character.CreateCharacter()
	} else {
		fmt.Println("Sauvegarde chargée avec succès.")
	}

	// lancer le menu
	menu.Run(hero)

	// à la sortie du menu → on sauvegarde
	if err := save.SaveCharacter(hero, "save.json"); err != nil {
		fmt.Println("Erreur lors de la sauvegarde :", err)
	} else {
		fmt.Println("Progression sauvegardée")
		time.Sleep(500 * time.Millisecond)
		intro.LoadingAnimation("Fermeture", 5)
	}
}
