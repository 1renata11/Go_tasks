package main

/* Срез justString ссылается на все 1024 байта строки v,
а не только на 100 байт, хоть и имеет доступ только к первым 100 байтам.
Это может привести к утечке памяти.
Решение - создать новую строку на основе только
нужных байтов*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func L115() {
	someFunc()
}
