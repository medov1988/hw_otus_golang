package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

var (
	re1 = regexp.MustCompile("[0-9!-,./:-@\\[-`\\{-~]|[\t]|[\n]| - ")
	re2 = regexp.MustCompile(" - ")
	re3 = regexp.MustCompile("([ ]{2,})")
)

func Top10(s string) []string {
	s = re1.ReplaceAllString(s, " ")
	s = re2.ReplaceAllString(s, " ")
	s = re3.ReplaceAllString(s, " ")
	s = strings.Trim(s, "- ")
	s = strings.ToLower(s)

	if len(s) == 0 {
		return []string{}
	}

	a := strings.Split(s, " ")
	m := map[string]int{}

	for _, e := range a {
		m[e]++
	}

	sortedStruct := make([]kv, len(m))
	i := 0
	for key, value := range m {
		sortedStruct[i] = kv{key, value}
		i++
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Value != sortedStruct[j].Value {
			return sortedStruct[i].Value > sortedStruct[j].Value
		}
		return sortedStruct[i].Key < sortedStruct[j].Key
	})

	result := make([]string, 0, 10)

	for k, v := range sortedStruct {
		result = append(result, v.Key)
		if k == 9 {
			break
		}
	}
	return result
}
