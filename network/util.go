package network

import (
	"errors"
	"strings"
)

func ReadComman(input string) (string, string, error) {
	t := strings.Split(input, "@")
	if len(t) != 2 {
		return "", "", errors.New("Invalid input string")
	}
	return t[0], t[1], nil
}
