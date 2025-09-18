package intro

import (
	"fmt"
	"time"
)

// Fonction utilitaire pour texte lent
func slowPrint(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

// Affiche le logo ASCII
func printAsciiLogo() {
	logo := "  _           _           _           \n" +
		" | |         | |         | |          \n" +
		" | |     __ _| |__  _   _| |__  _   _ \n" +
		" | |    / _' | '_ \\| | | | '_ \\| | | |\n" +
		" | |___| (_| | |_) | |_| | |_) | |_| |\n" +
		" |______\\__,_|_.__/ \\__,_|_.__/ \\__,_|\n" +
		"  _______    _ _                      \n" +
		" |__   __|  (_) |                     \n" +
		"    | | __ _ _| |                     \n" +
		"    | |/ _' | | |                     \n" +
		"    | | (_| | | |                     \n" +
		"    |_|\\__,_|_|_|                     \n" +
		"\n\n" +
		"             [Labubu Tails] — Le sourire creux s'etend...\n"
	fmt.Println(logo)
}

// Animation de chargement
func LoadingAnimation(text string, cycles int) {
	for i := 0; i < cycles; i++ {
		for _, c := range []string{"|", "/", "-", "\\"} {
			fmt.Printf("\r%s %s", text, c)
			time.Sleep(150 * time.Millisecond)
		}
	}
	fmt.Print("\r") // Clear line
}

// Fonction principale de l’intro
func RunIntro() {
	printAsciiLogo()
	time.Sleep(1 * time.Second)

	slowPrint("Bienvenue, voyageur dans l'ombre des Labubus...", 50*time.Millisecond)
	time.Sleep(1 * time.Second)

	slowPrint("Connexion au Nexus des Labubus...", 40*time.Millisecond)
	LoadingAnimation("Chargement", 10)

	slowPrint("\nLe silence tremble... Les cauchemars s'eveillent...", 50*time.Millisecond)
	time.Sleep(1 * time.Second)

	slowPrint("\n[Appuyez sur Entrée pour continuer]", 40*time.Millisecond)
	fmt.Scanln()
}
