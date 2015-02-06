package main

import (
	"fmt"
	"github.com/james-maloney/termstyle"
)

func main() {
	c := 0
	for i := 0; i < 16; i++ {
		f := ""
		for j := 0; j < 16; j++ {
			f += termstyle.FG[c] + fmt.Sprint(c) + termstyle.C + " "
			c++
		}
		fmt.Println(f)
	}

	c = 0
	for i := 0; i < 16; i++ {
		f := ""
		for j := 0; j < 16; j++ {
			f += termstyle.BG[c] + fmt.Sprint(c) + termstyle.C + " "
			c++
		}
		fmt.Println(f)
	}
}
