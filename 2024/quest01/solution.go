package main

import (
  "everybody-codes/utils"
)

const (
  inputPart1 = "2024/inputs/quest01/part1.txt"
  inputPart2 = "2024/inputs/quest01/part2.txt"
  inputPart3 = "2024/inputs/quest01/part3.txt"
)

func main() {
  utils.Run("Part 1", part1)
  utils.Run("Part 2", part2)
  utils.Run("Part 3", part3)
}

func part1() int {
  creatures, _ := parseInput(inputPart1)

  totalPotions := 0
  for _, creature := range creatures {
    totalPotions += getPotionCost(creature)
  }
  return totalPotions
}

func part2() int {
  creatures, _ := parseInput(inputPart2)

  totalPotions := 0
  // Process in pairs (every 2 characters)
  for i := 0; i < len(creatures); i += 2 {
    // Extract the pair
    pair := creatures[i : i+2]
    totalPotions += calculatePairCost(pair)
  }

  return totalPotions
}

func part3() int {
  creatures, _ := parseInput(inputPart3)

  totalPotions := 0
  for i := 0; i < len(creatures); i += 3 {
    // Extract the pair
    trio := creatures[i : i+3]
    totalPotions += calculateTrioCost(trio)
  }

  return totalPotions
}
