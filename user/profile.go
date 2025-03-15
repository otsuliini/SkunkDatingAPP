package account

type raccoon struct {
	// Name of the raccoon
	Name string
	// Age of the raccoon
	Age int
	// Rizz
	Rizz string
	//Aura
	Aura int
}

func CreateNewRaccoon(name string, age int, rizz string, aura int) *raccoon {
	return &raccoon{Name: name, Age: age, Rizz: "rizz", Aura: aura}
}
