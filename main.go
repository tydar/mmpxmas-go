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

	apiKey, keyYes := os.LookupEnv("XMAS_COINMARKETCAP_KEY")

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
	case "rows":
		rows()
	case "mod2":
		evenOdd()
	case "alt":
		redGreen()
	case "btc":
		if keyYes {
			btcMinute(apiKey)
		} else {
			fmt.Println("set environment variable XMAS_COINMARKETCAP_KEY")
		}
	default:
		fmt.Println("supported subcommands: 'ants', 'clear', 'rows', 'mod2', 'alt', 'btc'")
	}
}

func ants() {
	for true {
		for i := range all {
			all[i].High()
			time.Sleep(time.Second / 5)
			all[i].Low()
		}
	}
}

func clear() {
	for i := range all {
		all[i].Low()
	}
}

func rows() {
	// light up each row one at a time
	for true {
		for i := range bottom {
			bottom[i].High()
		}

		time.Sleep(time.Second / 3)

		for i := range bottom {
			bottom[i].Low()
		}

		for i := range middle {
			middle[i].High()
		}

		time.Sleep(time.Second / 3)

		for i := range middle {
			middle[i].Low()
		}

		for i := range top {
			top[i].High()
		}

		time.Sleep(time.Second / 3)

		for i := range top {
			top[i].Low()
		}

		star.High()
		time.Sleep(time.Second / 3)
		star.Low()
	}
}

func evenOdd() {
	// light even LEDs then odd LEDs
	// star stays lit
	star.High()

	for true {
		for i := range all {
			if i%2 == 0 && i != 10 {
				all[i].High()
				time.Sleep(time.Second / 5)
			}
		}

		for i := range all {
			if i%2 == 0 && i != 10 {
				all[i].Low()
			}
		}

		for i := range all {
			if i%2 != 0 {
				all[i].High()
				time.Sleep(time.Second / 5)
			}
		}

		for i := range all {
			if i%2 != 0 {
				all[i].Low()
			}
		}
	}
}

func redGreen() {
	red := []int{1, 3, 4, 6, 8}
	green := []int{0, 2, 5, 7, 9}

	star.High()
	for true {
		for _, n := range red {
			all[n].High()
		}
		time.Sleep(time.Second / 3)
		for _, n := range red {
			all[n].Low()
		}

		for _, n := range green {
			all[n].High()
		}
		time.Sleep(time.Second / 3)
		for _, n := range green {
			all[n].Low()
		}
	}
}
