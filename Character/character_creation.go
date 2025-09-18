package character

import (
	"Projet-Red_GoonLegend/classes"
	"Projet-Red_GoonLegend/items"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

// CreateCharacter : demande à l’utilisateur de créer son personnage
func CreateCharacter() *Character {
	reader := bufio.NewReader(os.Stdin)

	// === Choix du nom ===
	var nom string
	re := regexp.MustCompile(`^[a-zA-Z ]+$`)
	for {
		fmt.Print("Entrez le nom de votre personnage (uniquement des lettres) : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if re.MatchString(input) {
			// Met en majuscule la 1ère lettre de chaque mot
			mots := strings.Fields(strings.ToLower(input))
			for i, mot := range mots {
				if len(mot) > 0 {
					mots[i] = strings.ToUpper(string(mot[0])) + mot[1:]
				}
			}
			nom = strings.Join(mots, " ")
			break
		} else {
			fmt.Println("Le nom doit contenir uniquement des lettres et des espaces. Réessayez.")
		}
	}

	// === Choix de la classe ===
	options := []string{}
	for _, c := range classes.Classes {
		options = append(options, fmt.Sprintf("%s (%d PV max)", c.Nom, c.PVMax))
	}

	prompt := promptui.Select{
		Label: "Choisissez une classe",
		Items: options,
		Size:  len(options),
	}
	index, _, _ := prompt.Run()
	classeChoisie := classes.Classes[index]

	// === Création du personnage avec InitCharacter ===
	heroVal := InitCharacter(
		nom,
		classeChoisie,
		[]items.Item{
			{Nom: "Potion de soin", Type: "Potion", Soin: 50},
			{Nom: "Potion de soin", Type: "Potion", Soin: 50},
			{Nom: "Potion de soin", Type: "Potion", Soin: 50},
		},
	)

	return &heroVal
}
