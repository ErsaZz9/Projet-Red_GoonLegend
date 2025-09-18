package menu

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/items"
	"Projet-Red_GoonLegend/utils"
	"fmt"

	"github.com/manifoldco/promptui"
)

// Menu principal
func Run(hero *character.Character) {
	for {
		utils.ClearScreen() // ⚡ nettoie l’écran avant chaque affichage

		options := []string{
			"Afficher les informations du personnage",
			"Accéder à l’inventaire",
			"Marchand",
			"Quitter",
		}

		prompt := promptui.Select{
			Label: "=== Menu Principal ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}", // ⚡ retire le "?"
			},
			Size: 8,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annulé.")
			return
		}

		switch index {
		case 0:
			utils.ClearScreen()
			hero.DisplayInfo()
			fmt.Println("\n(Appuyez sur Entrée pour continuer)")
			fmt.Scanln()
		case 1:
			accessInventoryMenu(hero)
		case 2:
			runShop(hero)
		case 3:
			fmt.Println("Fermeture")
			return
		}
	}
}

// Sous-menu inventaire
func accessInventoryMenu(hero *character.Character) {
	for {
		utils.ClearScreen() // ⚡ efface avant de réafficher l’inventaire

		options := make([]string, 0, len(hero.Inventaire)+1)
		for _, it := range hero.Inventaire {
			options = append(options, it.Nom)
		}
		options = append(options, "← Retour")

		prompt := promptui.Select{
			Label: "=== Inventaire ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 10,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annulé.")
			return
		}

		// Retour
		if index == len(options)-1 {
			return
		}

		item := hero.Inventaire[index]

		switch item.Nom {
		case "Potion de soin":
			hero.UsePotion()
		case "Potion de poison":
			hero.ApplyPoison(item.Degats, 3)
			hero.RemoveInventory(item.Nom)
		case "Livre de Sort : Boule de Feu":
			hero.SpellBook("Boule de Feu")
			hero.RemoveInventory(item.Nom)
			fmt.Println("🔥 Vous avez appris : Boule de Feu !")
		default:
			fmt.Printf("Impossible d’utiliser %s pour l’instant.\n", item.Nom)
		}

		fmt.Println("\n(Appuyez sur Entrée pour continuer)")
		fmt.Scanln()
	}
}

// Sous-menu marchand
func runShop(hero *character.Character) {
	for {
		utils.ClearScreen() // ⚡ efface avant d’afficher le marchand

		options := []string{
			"Acheter une potion de soin (gratuite)",
			"Acheter une potion de poison (gratuite)",
			"Acheter un Livre de Sort : Boule de Feu (gratuit)",
			"← Retour",
		}

		prompt := promptui.Select{
			Label: "=== Marchand ===",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label: "{{ . }}",
			},
			Size: 8,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Annulé.")
			return
		}

		switch index {
		case 0:
			if hero.IsInventoryFull() {
				fmt.Println("Vous ne pouvez pas stocker plus de 10 objets dans votre inventaire !")
			} else {
				hero.AddInventory(items.Item{Nom: "Potion de soin", Type: "Potion", Soin: 50})
				fmt.Println("Vous avez acheté : Potion de soin")
			}
		case 1:
			if hero.IsInventoryFull() {
				fmt.Println("Vous ne pouvez pas stocker plus de 10 objets dans votre inventaire !")
			} else {
				hero.AddInventory(items.Item{Nom: "Potion de poison", Type: "Potion", Degats: 10})
				fmt.Println("Vous avez acheté : Potion de poison")
			}
		case 2:
			if hero.IsInventoryFull() {
				fmt.Println("Vous ne pouvez pas stocker plus de 10 objets dans votre inventaire !")
			} else {
				hero.AddInventory(items.Item{Nom: "Livre de Sort : Boule de Feu", Type: "Livre"})
				fmt.Println("Vous avez acheté : Livre de Sort : Boule de Feu")
			}
		case 3:
			return
		}

		fmt.Println("\n(Appuyez sur Entrée pour continuer)")
		fmt.Scanln()
	}
}
