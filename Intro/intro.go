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
	logo := `
	░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓████████▓▒░░▒▓███████▓▒░             ░▒▓██████▓▒░░▒▓████████▓▒░      
	░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             
	░▒▓█▓▒░     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             
	░▒▓██████▓▒░░▒▓█▓▒░      ░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░  ░▒▓██████▓▒░             ░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░        
	░▒▓█▓▒░     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             ░▒▓█▓▒░            ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             
	░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             ░▒▓█▓▒░            ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░             
	░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓████████▓▒░▒▓███████▓▒░              ░▒▓██████▓▒░░▒▓█▓▒░             
                                                                                                                           
                                                                                                                           
	░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░            ░▒▓███████▓▒░ ░▒▓██████▓▒░▒▓████████▓▒░                                
	   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░                                    
	   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░                                    
	   ░▒▓█▓▒░   ░▒▓████████▓▒░▒▓██████▓▒░              ░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░                                    
	   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░                                    
	   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░                   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░                                    
	   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░            ░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░  ░▒▓█▓▒░                                    
                                                                                                                           
                                                                                                                           
                                                                                                                           
             [ECHOES*OF*THE*ROT] — The Synaptic Decay Begins...
    ` // Violet/Magenta
	fmt.Println(logo) // Reset
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

	slowPrint("Bienvenue, Porteur de la volonté rot...", 50*time.Millisecond)
	time.Sleep(1 * time.Second)

	slowPrint("Connexion au monde de Rotborne...", 40*time.Millisecond)
	LoadingAnimation("Chargement", 10)

	slowPrint("\nLe voile s'efface... Les souvenirs reviennent...", 50*time.Millisecond)
	time.Sleep(1 * time.Second)

	slowPrint("\n[Appuyez sur Entrée pour continuer]", 40*time.Millisecond)
	fmt.Scanln()
}
