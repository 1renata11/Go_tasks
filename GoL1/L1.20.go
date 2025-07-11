package main

func TurnToMas(s string) []string {
	var mas []string
	temp := ""
	for _, x := range s {
		if x == ' ' {
			mas = append(mas, temp)
			temp = ""
		} else {
			temp += string(x)
		}
	}
	mas = append(mas, temp)
	return mas
}

func ReverseSentence(mas []string) []string {
	for i, j := 0, len(mas)-1; i < j; i, j = i+1, j-1 {
		mas[i], mas[j] = mas[j], mas[i]
	}
	return mas
}

func TurnToString(s []string) string {
	temp := ""
	for i := 0; i < len(s)-1; i++ {
		temp += s[i] + " "
	}
	temp += s[len(s)-1]
	return temp
}

func L120() {
	s := "абвгд еёжзи йклмн"
	print(TurnToString(ReverseSentence(TurnToMas(s))))

}
