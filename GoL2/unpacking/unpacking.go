package unpacking

import (
	"errors"
	"fmt"
	"unicode"
)

//Unpacking распаковывает строку
func Unpacking(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	var result string
	var before rune
	num := 0
	started := false
	escaped := false

	for _, r := range s {
		if escaped {
			if started {
				repeat := num
				if repeat == 0 {
					repeat = 1
				}
				for i := 0; i < repeat; i++ {
					result += string(before)
				}
			}
			before = r
			num = 0
			started = true
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if !started {
				return "", errors.New("Некорректная строка: число без символа")
			}
			num = num*10 + int(r-'0')
		} else {
			if started {
				repeat := num
				if repeat == 0 {
					repeat = 1
				}
				for i := 0; i < repeat; i++ {
					result += string(before)
				}
			}
			before = r
			num = 0
			started = true
		}
	}

	if escaped {
		return "", errors.New("Некорректная строка: символ экранирования без символа")
	}

	if started {
		repeat := num
		if repeat == 0 {
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			result += string(before)
		}
	}

	return result, nil
}

func TestSimple() {
	s := "a5"
	result, er := Unpacking(s)
	if er != nil {
		fmt.Println("Ошибка в функции")
		return
	}
	if result != "aaaaa" {
		fmt.Println("Несовпадение с ответом")
		return
	}
	fmt.Println("Корректно")
}

func TestComp() {
	s := "a10"
	result, er := Unpacking(s)
	if er != nil {
		fmt.Println("Ошибка в функции")
		return
	}
	if result != "aaaaaaaaaa" {
		fmt.Println("Несовпадение с ответом")
		return
	}
	fmt.Println("Корректно")
}

func TestEscaped() {
	s := "qwe\\45"
	result, er := Unpacking(s)
	if er != nil {
		fmt.Println("Ошибка в функции")
		return
	}
	if result != "qwe44444" {
		fmt.Println("Несовпадение с ответом")
		return
	}
	fmt.Println("Корректно")
}

func TestEmpty() {
	s := ""
	result, er := Unpacking(s)
	if er != nil {
		fmt.Println("Ошибка в функции")
		return
	}
	if result != "" {
		fmt.Println("Несовпадение с ответом")
		return
	}
	fmt.Println("Корректно")
}

func TestUncorrect() {
	s := "45"
	_, er := Unpacking(s)
	if er != nil {
		fmt.Println("Корректно")
	}
}

func TestAll() {
	TestSimple()
	TestComp()
	TestEscaped()
	TestEmpty()
	TestUncorrect()
}
