package main

import (
    "fmt"
    "time"
)

func slowPrint(text string, delay time.Duration) {
    for _, char := range text {
        fmt.Printf("%c", char)
        time.Sleep(delay)
    }
    fmt.Println()
}

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
          

             [ROT★BORNE] — The Synaptic Decay Begins...
    `
    fmt.Println("\033[35m")
    fmt.Println(logo)
    fmt.Print("\033[0m") // reset
}

func loadingAnimation(text string, cycles int) {
    for i := 0; i < cycles; i++ {
        for _, c := range []string{"|", "/", "-", "\\"} {
            fmt.Printf("\r%s %s", text, c)
            time.Sleep(150 * time.Millisecond)
        }
    }
    fmt.Print("\r") // Clear line
}

func main() {
    printAsciiLogo()
    time.Sleep(1 * time.Second)

    slowPrint("Bienvenue, Porteur de la Pourriture Mentale...", 50*time.Millisecond)
    time.Sleep(1 * time.Second)

    slowPrint("Connexion au monde de Rotborne...", 40*time.Millisecond)
    loadingAnimation("Chargement", 10)

    slowPrint("\nLe voile s'efface... Les souvenirs reviennent...", 50*time.Millisecond)
    time.Sleep(1 * time.Second)

    slowPrint("\n[Appuyez sur Entrée pour continuer]", 40*time.Millisecond)
    fmt.Scanln() // pause jusqu’à appui sur entrée
}