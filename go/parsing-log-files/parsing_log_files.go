package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	exp := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return exp.MatchString(text)
}

func SplitLogLine(text string) []string {
	exp := regexp.MustCompile(`<[-=~*]*>`)
	return exp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	exp := regexp.MustCompile("(?i)\".*password.*\"")
	var count int
	for _, val := range lines {
		if exp.MatchString(val) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	exp := regexp.MustCompile(`end-of-line\d*`)
	return exp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	exp := regexp.MustCompile(`User\s+(\w+)`)
	var res []string
	for _, val := range lines {
		matches := exp.FindStringSubmatch(val)
		if matches != nil {
			val = fmt.Sprintf("[USR] %s %s", matches[1], val)
		}
		res = append(res, val)
	}
	return res
}
