package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	after        int
	before       int
	around       int
	number       bool
	ignore       bool
	reverse      bool
	fix          bool
	numberBefore bool
	pattern      string
	file         string
)

//Функция, читающая строки из STDIN или файла
func ReadLines(files *string, lines *[]string) {
	switch len(*files) {
	case 0:
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			*lines = append(*lines, scanner.Text())
		}
	default:
		file, err := os.Open(*files)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при открытии файла")
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			*lines = append(*lines, scanner.Text())
		}
	}
}

//Функция, читающая аргументы из команды
func ReadArgs() {
	A := flag.Int("A", 0, "вывести контекст после")
	B := flag.Int("B", 0, "вывеати контекст перед")
	C := flag.Int("C", 0, "вывести контекст вокруг")
	flag.BoolVar(&number, "c", false, "количество строк")
	flag.BoolVar(&ignore, "i", false, "игнорировать регистр")
	flag.BoolVar(&reverse, "v", false, "инвертировать")
	flag.BoolVar(&fix, "F", false, "востпринимать шаблон как фиксированную строку")
	flag.BoolVar(&numberBefore, "n", false, "номер строки перед найденной строкой")
	flag.Parse()
	if *C > 0 {
		*A, *B = *C, *C
	}
	after = *A
	before = *B
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(os.Stderr, "Нет паттерна")
		os.Exit(2)
	}
	pattern = args[0]
	if len(args) > 1 {
		file = args[1]
	}

}

type Matcher interface {
	Match(string) bool
}

type Regular struct {
	reg *regexp.Regexp
}

func (reg Regular) Match(s string) bool {
	return reg.reg.MatchString(s)
}

type Fixed struct {
	str string
	ign bool
}

func (fixed Fixed) Match(s string) bool {
	if fixed.ign {
		return strings.Contains(strings.ToLower(s), strings.ToLower(fixed.str))
	}
	return strings.Contains(s, fixed.str)
}

//Функция, которая проверяет строку или регулярное выражение на совпадение
func BuildMatcher() Matcher {
	if fix {
		return Fixed{pattern, ignore}
	}
	pat := pattern
	if ignore {
		pat = "(?i)" + pat
	}
	regular := regexp.MustCompile(pat)
	return Regular{regular}
}

//Функция, печатающая строку и номер строки перед ней, если выставлен флаг numberBefore
func printLine(num int, line string) {
	if numberBefore {
		fmt.Printf("%d:%s\n", num, line)
	} else {
		fmt.Println(line)
	}
}

//Функция grep
func grep(m Matcher, lines *[]string) {
	matchCount := 0

	for i, line := range *lines {
		match := m.Match(line)
		if reverse {
			match = !match
		}
		if match {
			matchCount++
			if !number {
				start := i - before
				if start < 0 {
					start = 0
				}
				for j := start; j <= i; j++ {
					printLine(j+1, (*lines)[j])
				}
				end := i + after
				if end >= len(*lines) {
					end = len(*lines) - 1
				}
				for j := i + 1; j <= end; j++ {
					printLine(j+1, (*lines)[j])
				}
			}
		}

	}
	if number {
		fmt.Println(matchCount)
	}
}

func main() {
	ReadArgs()
	var lines []string
	ReadLines(&file, &lines)
	m := BuildMatcher()
	grep(m, &lines)
}
