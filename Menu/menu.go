package menu

import (
	"Projet-Red_GoonLegend/character"
	"Projet-Red_GoonLegend/items"
	"fmt"
)

// je vais pas expliquer ça qd mm
func Run(hero *character.Character) {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")
		fmt.Print("Votre choix : ")

		// on crée une variable auquel on va attribuer un numéro via un scan ln (une input quoi)
		var choix int
		fmt.Scanln(&choix)

		// les switch case 👅

		switch choix {
		case 1:
			hero.DisplayInfo()
		case 2:
			accessInventoryMenu(hero)
		case 3:
			runShop(hero)
		case 4:
			fmt.Println("Bye")
			return
			// si l'utiliseur met autre chose quz le nombre demandé, 🖕
		default:
			fmt.Println("Vilain que tu es ! Tu ne sais pas compter jusqu'à 4 ? Réésaye")
		}
	}
}

// Menu après (case 2) pour l'inventaire
func accessInventoryMenu(hero *character.Character) {
	for {
		hero.AccessInventory()
		fmt.Println("1. Utiliser une potion")
		fmt.Println("2. Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			items.TakePot(hero)
		case 2:
			return
		default:
			fmt.Println("🖕")
		}
	}
}

func runShop(hero *character.Character) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Println("1. Acheter une potion de vie (gratuitement)")
		fmt.Println("2. Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			hero.AddInventory("Potion")
			fmt.Println("🛒 Vous avez acheté : Potion")
		case 2:
			return
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}
	}
}
