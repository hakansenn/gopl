package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567891"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	buf := new(bytes.Buffer)
	// Write first group of numbers
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	// The rest
	for j := i + 3; j < n;  {
		buf.WriteString("," + s[i:j])
		i, j = j, j+3
	}
	buf.WriteString("," + s[i:])
	return buf.String()
}
