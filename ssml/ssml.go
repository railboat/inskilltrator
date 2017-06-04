// Package ssml contains methods useful for working with SSML functions.
package ssml

import (
	"bytes"
	"fmt"
	"time"
)

// Break returns a series of one or more <break> ssml tags to get the required length of pause.
func Break(t time.Duration) string {
	buf := bytes.Buffer{}
	for ; t > 10*time.Second; t -= 10 * time.Second {
		buf.WriteString(fmt.Sprintf(`<break time="%s"/>`, 10*time.Second))
	}
	buf.WriteString(fmt.Sprintf(`<break time="%s"/>`, t))
	return buf.String()
}

// Interjection returns the SSML tag for an interjection around a string.
func Interjection(interjection string) string {
	return fmt.Sprintf(`<say-as interpret-as="interjection">%s</say-as>`, interjection)
}
