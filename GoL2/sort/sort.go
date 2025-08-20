package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	numeric bool
	reverse bool
	unique  bool
	key     bool
	keyNum  int
	files   []string
)

const chunkSize int = 100000

func main() {
	args := os.Args[1:]
	ReadArgs(&args)
	var lines []string
	ReadLines(&files, &lines)
	Sort(&lines)
	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при создании файла")
		os.Exit(1)
	}
	defer output.Close()
	writer := bufio.NewWriter(output)
	defer writer.Flush()
	for _, vr := range lines {
		writer.WriteString(vr + "\n")
	}
}

func ReadArgs(args *[]string) {
	for i := 0; i < len(*args); i++ {
		arg := (*args)[i]
		if arg == "--" {
			i++
			for ; i < len(*args); i++ {
				files = append(files, (*args)[i])
			}
			break
		}
		if arg[0] == '-' && len(arg) > 0 {
			if arg == "-k" {
				key = true
				i++
				if i >= len(*args) {
					fmt.Fprintln(os.Stderr, "Флаг -k требует аргумент")
					os.Exit(1)
				}
				var err error
				keyNum, err = strconv.Atoi((*args)[i])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Неверный аргумент для -k, должно быть число")
					os.Exit(1)
				}
			}
			for _, ch := range arg {
				switch ch {
				case 'n':
					numeric = true
				case 'r':
					reverse = true
				case 'u':
					unique = true
				case 'k':
					key = true
				case '-':
					continue
				default:
					fmt.Fprintln(os.Stderr, "Неизвестный флаг")
					os.Exit(1)
				}
			}
		} else {
			files = append(files, arg)
		}
	}
}

func ReadLines(files *[]string, lines *[]string) {
	switch len(*files) {
	case 0:
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			*lines = append(*lines, scanner.Text())
		}
	default:
		for i := 0; i < len(*files); i++ {
			file, err := os.Open((*files)[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при открытии файла")
				continue
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				*lines = append(*lines, scanner.Text())
			}
		}
	}
}

func Sort(lines *[]string) {
	if unique && len(*lines) > 0 {
		var uniqueLines []string
		if key {
			m := make(map[string]string)
			for i, _ := range *lines {
				str := strings.Fields((*lines)[i])
				if len(str) > keyNum {
					fmt.Fprintf(os.Stderr, "Нет столбца под номером %d", keyNum)
					os.Exit(1)
				}
				m[str[keyNum-1]] = (*lines)[i]
			}
			for _, vr := range m {
				uniqueLines = append(uniqueLines, vr)
			}
		} else {
			m := make(map[string]int)
			for _, v := range *lines {
				m[v] += 1
			}
			for key, _ := range m {
				uniqueLines = append(uniqueLines, key)
			}
		}
		*lines = uniqueLines
	}
	if numeric {
		if key {
			sort.Slice(*lines, func(i, j int) bool {
				var a, b float64
				str1 := strings.Fields((*lines)[i])
				str2 := strings.Fields((*lines)[j])
				if len(str1) > keyNum || len(str2) > keyNum {
					fmt.Fprintf(os.Stderr, "Нет столбца под номером %d", keyNum)
					os.Exit(1)
				}
				a, errA := strconv.ParseFloat(str1[keyNum-1], 64)
				b, errB := strconv.ParseFloat(str2[keyNum-1], 64)
				if errA != nil || errB != nil {
					return str1[keyNum-1] < str2[keyNum-1]
				}
				return a < b
			})
		} else {
			sort.Slice(*lines, func(i, j int) bool {
				a, errA := strconv.ParseFloat((*lines)[i], 64)
				b, errB := strconv.ParseFloat((*lines)[j], 64)
				if errA != nil || errB != nil {
					return (*lines)[i] < (*lines)[j]
				}
				return a < b
			})
		}
	} else {
		if key {
			sort.Slice(*lines, func(i, j int) bool {
				str1 := strings.Fields((*lines)[i])
				str2 := strings.Fields((*lines)[j])
				if len(str1) > keyNum || len(str2) > keyNum {
					fmt.Fprintf(os.Stderr, "Нет столбца под номером %c", keyNum)
					os.Exit(1)
				}
				return str1[keyNum-1] < str2[keyNum-1]
			})
		} else {
			sort.Strings(*lines)
		}
	}
	if reverse {
		for i, j := 0, len(*lines)-1; i <= j; i, j = i+1, j-1 {
			(*lines)[i], (*lines)[j] = (*lines)[j], (*lines)[i]
		}
	}

}
