package algorithms

// As people come to the dance floor, they should be paired off as quickly as possible: man with woman, man with woman, all the way down the line. If several men arrive in a row, they should be paired in the order they came, and likewise if several women do. Maintain a queue of "spares" (men for whom you have no women yet, or vice versa), and pair them as appropriate.
// Input:
// []Person{
// 		John,
// 		Jane,
// 		Serena,
// 		Luna,
// 		Darien,
// 		Artemis,
// 		Rei,
// 		Nicolas,
// 	}
// Output:
// [][]Person{
// 		{John, Jane},
// 		{Serena, Darien},
// 		{Luna, Artemis},
// 		{Rei, Nicolas},
// 	}



type Person struct{
	Name 		string
	Gender 	string
}

const (
	MAN = "man"
	WOMAN = "woman"
)

var John Person = Person{Name: "John", Gender: MAN}
var Jane Person = Person{Name: "Jane", Gender: WOMAN}
var Serena Person = Person{Name: "Serena", Gender: WOMAN}
var Luna Person = Person{Name: "Luna", Gender: WOMAN}
var Darien Person = Person{Name: "Darien", Gender: MAN}
var Artemis Person = Person{Name: "Artemis", Gender: MAN}
var Rei Person = Person{Name: "Rei", Gender: WOMAN}
var Nicolas Person = Person{Name: "Nicolás", Gender: MAN}

func SquareDancePairing([]Person) [][]Person{
	return [][]Person{}
}