package twofer

import (
	"fmt"
)

func ShareWith(name string) string {
	receiver := name
	if receiver == "" {
		receiver = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", receiver)
}
