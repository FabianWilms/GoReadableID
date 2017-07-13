package readableId

import (
	"time"
	"math/rand"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))
var adjectives_length = len(adjectives)
var animals_length = len(animals)

func GetRandomID() string {

	var randAnimal, randAdjectiveA, randAdjectiveB int

	randAnimal = r.Intn(animals_length)
	randAdjectiveA = r.Intn(adjectives_length)
	randAdjectiveB = r.Intn(adjectives_length)

	return adjectives[randAdjectiveA] + "-" + adjectives[randAdjectiveB] + "-" + animals[randAnimal]

}