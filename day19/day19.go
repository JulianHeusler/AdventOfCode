package day19

import "adventofcode/util"

type Blueprint struct {
	robots []RobotRecipe
}

type RobotRecipe struct {
	produces     string // maybe add enum?
	oreCost      int
	clayCost     int
	obsidianCost int
}

func Solve(lines []string) (part1 int, part2 int) {
	parse(lines)
	return 0, 0
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
