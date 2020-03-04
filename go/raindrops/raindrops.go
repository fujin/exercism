package raindrops

import (
	"bytes"
	"fmt"
)

// Convert function receives an integer describing a sound and appends to a buffer based on the sound. It's based on the 'fizz buzz' concept.
func Convert(sound int) string {
	var buffer bytes.Buffer
	if sound%3 == 0 {
		buffer.WriteString("Pling")
	}

	if sound%5 == 0 {
		buffer.WriteString("Plang")
	}

	if sound%7 == 0 {
		buffer.WriteString("Plong")
	}

	if sound%3 != 0 &&
		sound%5 != 0 &&
		sound%7 != 0 {
		buffer.WriteString(fmt.Sprintf("%d", sound))
	}
	return buffer.String()
}
