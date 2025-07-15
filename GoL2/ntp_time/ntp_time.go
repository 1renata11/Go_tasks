package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

//CurrentTime возвращает текущее время с NTP-сервера
func CurrentTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return time
}

func main() {
	fmt.Println(CurrentTime())
}
