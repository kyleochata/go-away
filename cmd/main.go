package main

import (
	"fmt"
	"time"

	"github.com/kyleochata/go-away/particles"
)

// func main() {
// 	coffee := particles.NewCoffee(7, 5)
// 	coffee.Start()

//		timer := time.NewTicker(time.Millisecond * 100.0)
//		for {
//			<-timer.C
//			fmt.Print("\033[H\033[2J") // clear the screen
//			coffee.Update()
//			fmt.Println(coffee.Display())
//		}
//	}
func main() {
	coffee := particles.NewCoffee(9, 5)
	coffee.Start()

	timer := time.NewTicker(time.Millisecond * 100)
	for {
		<-timer.C
		fmt.Print("\033[H") // move cursor to the top without clearing the screen
		coffee.Update()
		fmt.Println(coffee.Display())
	}
}
