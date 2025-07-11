package main

import "fmt"

/* Клиент ожидает NewPrinter
У нас есть MyOldPrinter, реализующий OldPrinter
Адаптер (PrinterAdapter) реализует NewPrinter, но внутри вызывает OldPrinter
Плюсы:
Позволяет использовать несовместимые интерфейсы
Упрощает миграцию на новые реализации
Минусы:
Может усложнить структуру кода*/

type OldPrinter interface {
	PrintOld(msg string) string
}

type MyOldPrinter struct{}

func (p *MyOldPrinter) PrintOld(msg string) string {
	return "Old: " + msg
}

type NewPrinter interface {
	Print(msg string)
}

type PrinterAdapter struct {
	OldPrinter OldPrinter
}

func (pa *PrinterAdapter) Print(msg string) {
	output := pa.OldPrinter.PrintOld(msg)
	fmt.Println(output)
}

func L121() {
	oldPrinter := &MyOldPrinter{}
	adapter := &PrinterAdapter{OldPrinter: oldPrinter}

	var printer NewPrinter = adapter
	printer.Print("Hi")
}
