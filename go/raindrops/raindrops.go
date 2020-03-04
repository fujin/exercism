package raindrops

import (
	"strconv"
)

// Convert function receives an integer describing a sound and appends to a buffer based on the sound. It's based on the 'fizz buzz' concept.
func Convert(sound int) (result string) {

	if sound%3 == 0 {
		result += "Pling"
	}

	if sound%5 == 0 {
		result += "Plang"
	}

	if sound%7 == 0 {
		result += "Plong"
	}

	if sound%3 != 0 &&
		sound%5 != 0 &&
		sound%7 != 0 {
		result += strconv.Itoa(sound)
	}
	return
}
