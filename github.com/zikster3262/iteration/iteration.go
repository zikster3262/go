package iteration

import "fmt"

func Repeat(a string, b int) string {
	var repeated string
	if b < 0 {
		b = 15
	}
	for i := 0; i < b; i++ {
		repeated += a
	}
	fmt.Println(repeated)
	return repeated
}
