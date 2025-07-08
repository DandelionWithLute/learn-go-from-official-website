package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
