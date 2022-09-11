// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

const stv string = "{ [ }"

func main() {
	// Convert to slice.
	c := strings.Split(stv, "")
	// String initial characters.
	sic := make([]string, 0)
	err := false

	for _, character := range c {
		// Validate open characters.
		if character == "{" || character == "(" || character == "[" {
			sic = append(sic, character)
			continue
		}

		// Validate close characters.
		if character == "}" || character == ")" || character == "]" {
			sic, err = removeCharacter(character, sic)
		}

		// In case that exist a close character we finish all operation.
		if err {
			fmt.Println(0)
			return
		}
	}

	// If the slice is empty then all character are correct.
	if len(sic) == 0 {
		fmt.Println(1)
		return
	}

	// If not, then some characters are incorrect.
	fmt.Println(0)
}

func removeCharacter(c string, sic []string) ([]string, bool) {
	// Character map.
	mc := make(map[string]string)
	mc["]"] = "["
	mc["}"] = "{"
	mc[")"] = "("

	// Character slice length.
	csl := len(sic)
	// If exist a final character we finish all operation.
	if csl == 0 {
		return sic, true
	}

	// Character in current position.
	cicp := sic[csl-1]
	// If the character is the opposite then remove the last character.
	if cicp == mc[c] {
		return sic[:csl-1], false
	}

	return sic, false
}
