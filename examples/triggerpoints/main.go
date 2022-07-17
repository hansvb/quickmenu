package main

import (
	"fmt"

	"github.com/hansvb/quickmenu"
)

func sayHi() {
	fmt.Println("Hi!")
}

func sayBye() {
	fmt.Println("Bye!")
}

func main() {
	// There are 4 trigger-points: Before/AfterMenu + Before/AfterExec
	// In this example we use them for styling the menu
	// but a more useful scenario might be to print some variables.

	// You can define functions for these trigger-points during initialization ...
	qm := quickmenu.QuickMenu{
		BeforeMenuFunc: func() {
			// Note that clearing the screen hides any entry-errors!
			fmt.Println("\033[H\033[2J") // clear screen (ANSI escape code)
			fmt.Println("=====================M=e=n=u==")
		},
		AfterMenuFunc: func() {
			fmt.Println("-----------------C-h-o-i-c-e--")
		},
	}

	qm.Add("Say Hi!", sayHi)
	qm.Add("Say Bye!", sayBye)

	// ... or define them later.
	qm.BeforeExecFunc = func() {
		fmt.Println("\n[start of execution]")
	}
	qm.AfterExecFunc = func() {
		fmt.Println("[end of execution]")
		fmt.Println("\nPress <enter> to continue...")
		fmt.Scanln()
	}

	qm.PromptLoop()
}
