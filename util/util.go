package util

import (
	"regexp"
	"strconv"
	"strings"
)

func FindFirstInt(input, regex string) int {
	expression := regexp.MustCompile(regex)
	return GetInt(expression.FindStringSubmatch(input)[1])
}

func FindIntSlice(input, regex, delimter string) (ints []int) {
	experssion := regexp.MustCompile(regex)
	items := strings.Split(experssion.FindStringSubmatch(input)[1], delimter)
	for _, item := range items {
		ints = append(ints, GetInt(item))
	}
	return ints
}

func FindStringSubmatch(input, regex string) []string {
	expression := regexp.MustCompile(regex)
	return expression.FindStringSubmatch(input)
}

func GetInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
