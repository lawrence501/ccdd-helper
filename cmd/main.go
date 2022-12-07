package main

import (
	"fmt"
	"math/rand"
	"time"

	"validate"

	"github.com/manifoldco/promptui"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for true {
		baseP := promptui.Prompt{
			Label:    "Command",
			Validate: validate.ValidateBase,
		}
		input, err := baseP.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println()
}
