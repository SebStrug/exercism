package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	age := seconds / 31557600
	switch planet {
	case "Mercury":
		age = age / 0.2408467
	case "Venus":
		age = age/ 0.61519726
	case "Mars":
		age = age/ 1.8808158
	case "Jupiter":
		age = age/ 11.862615
	case "Saturn":
		age = age/ 29.447498
	case "Uranus":
		age = age/ 84.016846
	case "Neptune":
		age = age/ 164.79132
	}
	return age
}

