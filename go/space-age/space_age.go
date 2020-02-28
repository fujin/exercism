package space

// Planet type represents the planet on which the birthday is being calculated for.
type Planet string

const earthOrbitalPeriod = 31557600

// multipliers are orbital periods compared to earth orbital periods
var planetOrbitalPeriod = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age function calculates the age (in years represented by float64) of a given entity, passed in seconds, for the chosen planet.
func Age(seconds float64, planet Planet) (years float64) {
	if planet == "Earth" {
		return (seconds / earthOrbitalPeriod)
	}

	orbitalPeriod, ok := planetOrbitalPeriod[planet]
	// If the map is missing an orbital period for the specified planet, return zero value.
	if !ok {
		return
	}

	return (seconds / (earthOrbitalPeriod * orbitalPeriod))
}
