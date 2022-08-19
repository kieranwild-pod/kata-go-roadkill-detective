package main

import (
	"bytes"
	"fmt"
	"strings"
)

var animals = []string{
	"aardvark", "alligator", "armadillo", "antelope", "baboon", "bear", "bobcat", "butterfly", "cat", "camel", "cow", "chameleon", "dog", "dolphin", "duck", "dragonfly", "eagle", "elephant", "emu", "echidna", "fish", "frog", "flamingo", "fox", "goat", "giraffe", "gibbon", "gecko", "hyena", "hippopotamus", "horse", "hamster", "insect", "impala", "iguana", "ibis", "jackal", "jaguar", "jellyfish", "kangaroo", "kiwi", "koala", "killerwhale", "lemur", "leopard", "llama", "lion", "monkey", "mouse", "moose", "meercat", "numbat", "newt", "ostrich", "otter", "octopus", "orangutan", "penguin", "panther", "parrot", "pig", "quail", "quokka", "quoll", "rat", "rhinoceros", "racoon", "reindeer", "rabbit", "snake", "squirrel", "sheep", "seal", "turtle", "tiger", "turkey", "tapir", "unicorn", "vampirebat", "vulture", "wombat", "walrus", "wildebeast", "wallaby", "yak", "zebra"}

func main() {
	var photo string
	fmt.Print("Enter roadkill photo: ")
	fmt.Scan(&photo)

	fmt.Println("Animal name is: ", roadkill(photo))
}

func roadkill(photo string) string {
	// sanitise photo - remove '=' & remove dupes
	var sanitisedPhoto = strings.Replace(photo, "=", "", -1)
	var noDupesPhoto = removeDups(sanitisedPhoto)

	for _, animalName := range animals {
		var noDupesAnimalName = removeDups(animalName)
		var reverseNoDupesPhoto = reverseString(noDupesPhoto)

		// match against photo
		if animalName == sanitisedPhoto {
			return animalName
		} else if noDupesAnimalName == noDupesPhoto {
			return animalName
		} else if noDupesAnimalName == reverseNoDupesPhoto {
			return animalName
		}
	}

	return "??"
}

func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}

func removeDups(s string) string {
	var buf bytes.Buffer
	var last rune
	for i, r := range s {
		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
		}
	}
	return buf.String()
}
