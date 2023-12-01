package day19

import (
	"adventofcode/util"
	"fmt"
	"math"
)

type Material int

const (
	Ore = iota
	Clay
	Obsidian
	Geode
)

type Blueprint struct {
	robots []RobotRecipe
}

type RobotRecipe struct {
	produces     string // maybe add enum?
	oreCost      int
	clayCost     int
	obsidianCost int
}

type Income struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type Wallet struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type State struct {
	income        Income
	wallet        Wallet
	remainingTime int
}

func Solve(lines []string) (part1 int, part2 int) {
	parse(lines)
	return 0, 0
}

func (m Material) EnumIndex() int {
	return int(m)
}

func solvePart1(blueprints []Blueprint) int {
	for _, blueprint := range blueprints {
		sim(blueprint, 24, Income{1, 0, 0, 0}, Wallet{0, 0, 0, 0})
	}
	return 0
}

func sim(b Blueprint, time int, income Income, wallet Wallet) int {
	var geodeCounts []int

	// for each robot recipe
	for i := 0; i < 4; i++ {
		robotRecipe := b.robots[i]
		recipeTime, err := timeToCompleteRecipe(robotRecipe, income, wallet)
		if err != nil {
			continue
		}
		remainingTime := time - recipeTime - 1
		if remainingTime < 0 {
			// TODO clac remaining stuff
		}
		geodeCounts = append(geodeCounts,
			sim(b, remainingTime, updateIncome(robotRecipe, income), updateWallet(robotRecipe, income, wallet, recipeTime)))
	}

	return maxOfIntSlice(geodeCounts)
}

func maxOfIntSlice(counts []int) (max int) {
	for _, c := range counts {
		if c > max {
			max = c
		}
	}
	return max
}

func updateIncome(r RobotRecipe, i Income) Income {
	switch r.produces {
	case "ore":
		i.ore++
		break
	case "clay":
		i.clay++
		break
	case "obsidian":
		i.obsidian++
		break
	case "geode":
		i.geode++
		break
	}
	return i
}

func updateWallet(r RobotRecipe, i Income, w Wallet, t int) Wallet {
	w.ore += (i.ore * t) - r.oreCost
	w.clay += (i.clay * t) - r.clayCost
	w.obsidian += (i.obsidian * t) - r.obsidianCost
	w.geode += (i.geode * t)
	return w
}

func timeToCompleteRecipe(recipe RobotRecipe, income Income, wallet Wallet) (time int, err error) {
	if income.ore <= 0 || income.clay <= 0 || income.obsidian <= 0 {
		return 0, fmt.Errorf("missing income robot")
	}

	ore := (wallet.ore - recipe.oreCost) / income.ore
	clay := (wallet.clay - recipe.clayCost) / income.clay
	obsidian := (wallet.obsidian - recipe.obsidianCost) / income.obsidian

	return int(math.Max(math.Max(float64(ore), float64(clay)), float64(obsidian))), nil
}

// keep in mind: blueprint index + 1
func parse(lines []string) (blueprints []Blueprint) {
	for _, line := range lines {
		oreRegex := util.FindStringSubmatch(line, `Each ore robot costs (\d+) ore.`)
		oreRobot := RobotRecipe{"ore", util.GetInt(oreRegex[1]), 0, 0}
		clayRegex := util.FindStringSubmatch(line, `Each clay robot costs (\d+) ore.`)
		clayRobot := RobotRecipe{"clay", util.GetInt(clayRegex[1]), 0, 0}
		obsidianRegex := util.FindStringSubmatch(line, `Each obsidian robot costs (\d+) ore and (\d+) clay.`)
		obsidianRobot := RobotRecipe{"obsidian", util.GetInt(obsidianRegex[1]), util.GetInt(obsidianRegex[2]), 0}
		geodeRegex := util.FindStringSubmatch(line, `Each geode robot costs (\d+) ore and (\d+) obsidian.`)
		geodeRobot := RobotRecipe{"geode", util.GetInt(geodeRegex[1]), 0, util.GetInt(geodeRegex[2])}
		blueprints = append(blueprints, Blueprint{[]RobotRecipe{oreRobot, clayRobot, obsidianRobot, geodeRobot}})
	}
	return blueprints
}
