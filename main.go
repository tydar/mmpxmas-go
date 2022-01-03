package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	botLeft     = rpio.Pin(11)
	botMidLeft  = rpio.Pin(16)
	botMidRight = rpio.Pin(17)
	botRight    = rpio.Pin(18)

	midLeft     = rpio.Pin(19)
	midMidLeft  = rpio.Pin(20)
	midMidRight = rpio.Pin(21)
	midRight    = rpio.Pin(22)

	topLeft  = rpio.Pin(23)
	topRight = rpio.Pin(24)

	star = rpio.Pin(12)

	bottom = []rpio.Pin{botLeft, botMidLeft, botMidRight, botRight}
	middle = []rpio.Pin{midLeft, midMidLeft, midMidRight, midRight}
	top    = []rpio.Pin{topLeft, topRight}

	all = []rpio.Pin{botLeft, botMidLeft, botMidRight, botRight,
		midLeft, midMidLeft, midMidRight, midRight,
		topLeft, topRight,
		star,
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	for i := range all {
		all[i].Output()
	}

	switch os.Args[1] {
	case "ants":
		ants()
	case "clear":
		clear()
	default:
		fmt.Println("supported subcommands: 'ants', 'clear'")
	}
}

func ants() {
	for true {
		for i := range all {
			all[i].Toggle()
			time.Sleep(time.Second / 2)
		}
	}
}

func clear() {
	for i := range all {
		all[i].Low()
	}
}
