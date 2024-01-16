package main

import "math/rand"

func CatchPokemon(experience int) bool {
  chance := rand.Intn(experience + 1)
  threshold := experience / 3
  return chance <= threshold
}

