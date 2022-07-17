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
	qm := quickmenu.QuickMenu{}
	qm.AfterMenuFunc = func() {
		fmt.Println("\nMake your choice:")
	}
	qm.Add("Say Hi!", sayHi)
	qm.Add("Say Bye!", sayBye)
	qm.PromptOnceAndQuit()
}
