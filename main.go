package main

import (
	"fmt"

	"github.com/aymanbagabas/go-osc52"
)

const shrug = `¯\_(ツ)_/¯`

func main() {
	fmt.Println(shrug)
	osc52.CopyPrimary(shrug)
	osc52.Copy(shrug)
}
