package main

import (
	"flag"
	"fmt"
	"github.com/crazy3lf/colorconv"
	"os"
)

func main() {

	red := uint(0)
	green := uint(0)
	blue := uint(0)
	h := float64(0)
	s := float64(0)
	v := float64(0)

	direction := "tohsv"
	flag.UintVar(&red, "r", 0, "Red value (decimal)")
	flag.UintVar(&green, "g", 0, "Green value (decimal)")
	flag.UintVar(&blue, "b", 0, "Blue value (decimal)")
	flag.StringVar(&direction, "direction", "tohsv", "Either tohsv OR torgb")
	flag.Float64Var(&h, "h", 0, "Hue")
	flag.Float64Var(&s, "s", 0, "Saturation")
	flag.Float64Var(&v, "v", 0, "Value")

	flag.Parse()

	if direction == "tohsv" {
		h, s, v := colorconv.RGBToHSV(uint8(red), uint8(green), uint8(blue))
		fmt.Printf("h=%f, s=%f, v=%f\n", h, s, v)
		return
	} else if direction == "torgb" {
		r, g, b, err := colorconv.HSVToRGB(h, s, v)
		if err != nil {
			fmt.Println("Error \n", err)
			os.Exit(-1)
		}
		fmt.Printf("r=%d, g=%d, b=%d\n", r, g, b)
	} else {
		fmt.Println("Error, must specify a direction, direction=torgb OR direction=tohsv")
	}
}
