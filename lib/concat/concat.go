// concat provides a highly optimized string concatenation method.
package concat

import "bytes"

// Takes an unlimited number of strings and concatenates them in the order
// that they were passed.
func Concat(strings ...string) string {
	var buffer bytes.Buffer
	for i := range strings {
		buffer.WriteString(strings[i])
	}

	return buffer.String()
}
