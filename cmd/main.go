package main

import (
	"fmt"
	"time"

	"github.com/kyleochata/go-away/particles"
)

func main() {
	coffee := particles.NewCoffee(9, 5)
	coffee.Start()

	timer := time.NewTicker(time.Millisecond * 100)
	for {
		<-timer.C
		fmt.Print("\033[H\033[J")
		coffee.Update()
		fmt.Println(coffee.Display())
	}
}
