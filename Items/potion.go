package items

import (
	"Projet-Red_GoonLegend/character"
	"fmt"
)

func TakePot(c *character.Character) {
	// regarde si une potion est dans l'inventaire ?
	potionIndex := -1
	for i, item := range c.Inventaire {
		if item == "Potion" {
			potionIndex = i
			break
		}
	}

	// -1 dans l'index n'étant pas compris, si aucune potion on va dessus et on print pas de potion
	if potionIndex == -1 {
		fmt.Println("Aucune potion de soin restante !")
		return
	}
	// on enleve la potion via append puisque on va l'a consomme, pour ça on garde tout ce qui est indexé avant et après
	c.Inventaire = append(c.Inventaire[:potionIndex], c.Inventaire[potionIndex+1:]...)

	// donc on soigne, si max bah pv = max

	c.PV += 50
	if c.PV > c.PVMax {
		c.PV = c.PVMax
	}
	// on print la confirmation + lesp pv actuels
	fmt.Printf("Potion utilisée ! PV actuels : %d/%d\n", c.PV, c.PVMax)
}
