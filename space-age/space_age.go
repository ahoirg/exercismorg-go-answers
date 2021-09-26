package space

type Planet string

const aEarthPeriotSeconds = float64(31557600)

func Age(second float64, planet Planet) float64 {
	return float64(howManyEarthPeriods(planet) * second / aEarthPeriotSeconds)
}

func howManyEarthPeriods(planet Planet) float64 {
	var periots float64
	switch planet {
	case "Mercury":
		periots = 1 / 0.2408467
	case "Venus":
		periots = 1 / 0.61519726
	case "Mars":
		periots = 1 / 1.8808158
	case "Jupiter":
		periots = 1 / 11.862615
	case "Saturn":
		periots = 1 / 29.447498
	case "Uranus":
		periots = 1 / 84.016846
	case "Neptune":
		periots = 1 / 164.79132
	default:
		periots = 1
	}
	return periots
}
