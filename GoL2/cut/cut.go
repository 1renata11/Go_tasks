package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	fieldsStr string
	delimiter string
	separated bool
	fields    []intRange
)

type intRange struct {
	start int
	end   int
}

//Функция, читающая строки из STDIN
func ReadLines(lines *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
}

//Функция, читающая аргументы из команды
func ReadArgs() {
	flag.StringVar(&fieldsStr, "f", "", "указание номеров полей")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки содержащие разделитель")
	flag.Parse()
}

//Функция, осуществляющая парсинг строки-аргумента флага -f
func ReadFields(s string) []intRange {
	if len(s) == 0 {
		fmt.Println(os.Stderr, "Пустая строка")
	}
	currentNum := 0
	currentIntRange := intRange{0, 0}
	var Range []intRange
	for _, char := range s {
		if unicode.IsDigit(char) {
			currentNum *= 10
			d, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(os.Stderr, "Не получилось сконвертировать строку в число")
				os.Exit(1)
			}
			currentNum += d
		}
		if char == ',' {
			currentIntRange.end = currentNum
			if currentIntRange.start == 0 {
				currentIntRange.start = currentNum
			}
			currentNum = 0
			Range = append(Range, currentIntRange)
			currentIntRange.start = 0
			currentIntRange.end = 0

		}
		if char == '-' {
			currentIntRange.start = currentNum
			currentNum = 0
		}
	}
	if s[len(s)-1] == '-' {
		fmt.Println(os.Stderr, "Некорректный диапазон")
		os.Exit(1)
	}
	currentIntRange.end = currentNum
	if currentIntRange.start == 0 {
		currentIntRange.start = currentNum
	}
	Range = append(Range, currentIntRange)

	return Range
}

//Функция cut
func cut(lines *[]string) {
	for _, str := range *lines {
		if separated && !strings.Contains(str, delimiter) {
			continue
		}
		cols := strings.Split(str, delimiter)
		out := []string{}
		for _, cur := range fields {
			for i := cur.start; i <= cur.end; i++ {
				if i-1 < len(cols) && i-1 >= 0 {
					out = append(out, cols[i-1])
				}
			}
		}
		if len(out) > 0 {
			fmt.Println(strings.Join(out, " "))
		}
	}
}

func main() {
	ReadArgs()
	var lines []string
	ReadLines(&lines)
	fields = ReadFields(fieldsStr)
	cut(&lines)
}
