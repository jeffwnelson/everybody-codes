package main

import (
  "os"
  "strings"
)

// parseInput reads the input file and returns the creature string
func parseInput(filename string) (string, error) {
  data, err := os.ReadFile(filename)
  if err != nil {
    return "", err
  }
  return strings.TrimSpace(string(data)), nil
}

// getPotionCost returns the number of potions needed for a given creature
func getPotionCost(creature rune) int {
  switch creature {
  case 'A':
    return 0
  case 'B':
    return 1
  case 'C':
    return 3
  case 'D':
    return 5
  case 'x':
    return 0
  default:
    return 0
  }
}

// calculatePairCost calculates the potion cost for a pair of creatures
func calculatePairCost(creaturePairs string) int {
  creatureA := rune(creaturePairs[0])
  creatureB := rune(creaturePairs[1])

  // Sum our creature pairs
  total := getPotionCost(creatureA) + getPotionCost(creatureB)

  // Add +2 if both are real creatures (pairing bonus of +1 each)
  // If a singel creature, then no buff added
  if creatureA != 'x' && creatureB != 'x' {
    total += 2
  }

  return total
}

// calculateTrioCost calculates the potion cost for a trio of creatures
func calculateTrioCost(creatureTrio string) int {
  creatureA := rune(creatureTrio[0])
  creatureB := rune(creatureTrio[1])
  creatureC := rune(creatureTrio[2])

  // Sum our creature trios
  total := getPotionCost(creatureA) + getPotionCost(creatureB) + getPotionCost(creatureC)

  // Count real creatures (not 'x')
  count := 0
  if creatureA != 'x' {
    count++
  }
  if creatureB != 'x' {
    count++
  }
  if creatureC != 'x' {
    count++
  }

  // Add bonus based on how many pair up
  // 3 real creatures: each pairs with 2 others = +2 per creature = +6 total
  // 2 real creatures: they pair with each other = +1 per creature = +2 total
  // 1 real creature: no pairing = +0

  // Sample Input
  // xBx = [0 + 1 + 0] = 1 + 0 buff (only a single creature)
  // AAA = [0 + 0 + 0] = 0 + 6 buff (3 real creatures, +2 each)
  // BCD = [1 + 3 + 5] = 9 + 6 buff (3 real creatures, +2 each)
  // xCC = [0 + 3 + 3] = 6 + 2 buff (2 real creaturues, +1 each)

  if count == 3 {
    total += 6
  } else if count == 2 {
    total += 2
  }

  return total
}
