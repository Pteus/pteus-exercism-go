package twofer

import "fmt"

// ShareWith receives the next persons name and prints the desired phrase
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
