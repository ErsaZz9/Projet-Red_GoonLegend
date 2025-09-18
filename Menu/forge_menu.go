package menu

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/forge"
	"Projet-Red_GoonLegend/utils"
	"fmt"

	"github.com/manifoldco/promptui"
)

func runForge(hero *character.Character) {
	for {
		utils.ClearScreen()

		// Construire menu dynamique depuis forge.ForgeRecipes
		options := make([]string, 0, len(forge.ForgeRecipes)+1)
		for _, r := range forge.ForgeRecipes {
			options = append(options, fmt.Sprintf("Fabriquer %s (%d or)", r.Nom, r.Cout))
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "=== Forgeron ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 8,
		}

		index, _, err := prompt.Run()
		if err != nil || index == len(options)-1 {
			return
		}

		choix := forge.ForgeRecipes[index]

		// Afficher les ressources nécessaires
		fmt.Println("\nRessources nécessaires :")
		for nom, qte := range choix.Materiaux {
			fmt.Printf("- %s x%d\n", nom, qte)
		}
		fmt.Println()

		// Vérifications séparées (chaque erreur affiche un message et fait un continue)
		if hero.Purse < choix.Cout {
			fmt.Println("Vous n'avez pas assez d'or !")
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		if hero.IsInventoryFull() {
			fmt.Println("Inventaire plein !")
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		if !hasMaterials(hero, choix.Materiaux) {
			fmt.Println("Vous n'avez pas les ressources nécessaires !")
			fmt.Println("\n(Appuyez sur Entree pour continuer)")
			fmt.Scanln()
			continue
		}

		// Si toutes les vérifications passent → fabrication
		hero.Purse -= choix.Cout
		consumeMaterials(hero, choix.Materiaux)
		hero.AddInventory(choix.Resultat)
		fmt.Printf("%s fabriqué ! Or restant : %d\n", choix.Nom, hero.Purse)

		fmt.Println("\n(Appuyez sur Entree pour continuer)")
		fmt.Scanln()
	}
}

// Vérifie si le joueur possède toutes les ressources
func hasMaterials(hero *character.Character, mats map[string]int) bool {
	counts := make(map[string]int)
	for _, it := range hero.Inventaire {
		counts[it.Nom]++
	}
	for nom, qte := range mats {
		if counts[nom] < qte {
			return false
		}
	}
	return true
}

// Retire les ressources utilisées de l'inventaire
func consumeMaterials(hero *character.Character, mats map[string]int) {
	for nom, qte := range mats {
		for i := 0; i < qte; i++ {
			hero.RemoveInventory(nom)
		}
	}
}
