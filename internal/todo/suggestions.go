package todo

import "strings"

func getLevenshteinDistance(s1, s2 string) int {
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)

	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	column := make([]int, len(s1)+1)

	for i := 1; i <= len(s1); i++ {
		column[i] = i
	}

	for i := 1; i <= len(s2); i++ {
		column[0] = i
		lastkey := i - 1
		for k := 1; k <= len(s1); k++ {
			oldkey := column[k]
			incr := 1
			if s1[k-1] == s2[i-1] {
				incr = 0
			}
			column[k] = min(column[k]+1, column[k-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}

	return column[len(s1)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func suggestCommand(input string) string {
	commands := []string{"add", "list", "edit", "done", "delete", "clear", "help", "exit", "quit"}

	closest := ""
	minDist := 3

	for _, cmd := range commands {
		dist := getLevenshteinDistance(input, cmd)
		if dist < minDist {
			minDist = dist
			closest = cmd
		}
	}

	return closest
}
