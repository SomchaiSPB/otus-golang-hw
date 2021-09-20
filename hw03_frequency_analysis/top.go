package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type FrequentWords struct {
	Key string
	Value int
}

type List []FrequentWords

func (l List) Less(i int, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Key < l[j].Key
	}
	return l[i].Value < l[j].Value
}

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) ToStringArray() []string {
	output := make([]string, 0)

	for _, val := range l {
		output = append(output, val.Key)
	}

	return output
}

func Top10(input string) []string {
	words := strings.Fields(input)
	tempMap := make(map[string]int)
    var	sortedResult []string

	for _, word := range words {
		tempMap[word] = tempMap[word] + 1
	}

	tempResult := rankWordTop(tempMap)
	tempIntArr := make(map[int][]string)

	for key, val := range tempResult {

	}

	return tempResult.ToStringArray()[:10]
}

func rankWordTop(dict map[string]int) List {
	result := make(List, len(dict))
	count := 0

	for key, word := range dict {
		result[count] = FrequentWords{key, word}
		count++
	}

	sort.Sort(sort.Reverse(result))

	return result
}
