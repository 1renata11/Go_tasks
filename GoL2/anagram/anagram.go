package main

import (
	"fmt"
	"sort"
)

//Функция, сортирующая строки
func sortRunes(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

//Функция, получающая на вход список строк, на выходе выводит словарь анаграмм
func anagram(list *[]string) map[string][]string {
	anagramMap := make(map[string][]string)
	for i := 0; i < len(*list); i++ {
		word := (*list)[i]
		word = sortRunes(word)
		anagramMap[word] = append(anagramMap[word], (*list)[i])
	}
	resultMap := make(map[string][]string)
	for _, value := range anagramMap {
		if len(value) == 1 {
			continue
		}
		resultMap[value[0]] = value
	}
	return resultMap

}

func main() {
	list1 := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	list2 := []string{}
	list3 := []string{"пятак", "листок"}
	fmt.Println(anagram(&list1))
	fmt.Println(anagram(&list2))
	fmt.Println(anagram(&list3))
}
