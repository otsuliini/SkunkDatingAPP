package user


type skunk struct {
	// Name of the raccoon
	Name string
	// Age of the raccoon
	Age int
	// Rizz
	Rizz string
	// Aura
	Aura int
}


func CreateNewRaccoon(name string, age int, rizz string, aura int) *skunk {
	return &skunk{Name: name, Age: age, Rizz: "rizz", Aura: aura}
}
