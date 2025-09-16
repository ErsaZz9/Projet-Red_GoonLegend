package character

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PV         int
	Inventaire []string
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pv int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PV:         pv,
		Inventaire: inventaire,
	}
}
