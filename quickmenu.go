package quickmenu

import (
	"fmt"
	"os"
	"strconv"
)

type QuickMenuItem struct {
	order       int
	description string
	function    func()
}

type QuickMenu struct {
	items          []QuickMenuItem
	BeforeMenuFunc func()
	AfterMenuFunc  func()
	BeforeExecFunc func()
	AfterExecFunc  func()
}

func (qm *QuickMenu) Add(description string, f func()) {
	n := QuickMenuItem{len(qm.items) + 1, description, f}
	qm.items = append(qm.items, n)
}

func (qm QuickMenu) PromptOnceAndQuit() {
	choice := 0

	if qm.BeforeMenuFunc != nil {
		qm.BeforeMenuFunc()
	}

	for _, qmi := range qm.items {
		fmt.Printf("%d. %s\n", qmi.order, qmi.description)
	}

	if qm.AfterMenuFunc != nil {
		qm.AfterMenuFunc()
	}

	var entry string
	n, err := fmt.Scanln(&entry)
	if n == 0 {
		// just exit on <enter>
		os.Exit(0)
	}
	if err != nil {
		fmt.Printf("Wrong input: [%s]\n", err)
		os.Exit(4)
	}

	choice, err = strconv.Atoi(entry)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	if choice > len(qm.items) || choice <= 0 {
		fmt.Printf("Choice [%d] not available\n", choice)
		os.Exit(2)
	}

	if qm.BeforeExecFunc != nil {
		qm.BeforeExecFunc()
	}

	for _, qmi := range qm.items {
		if choice == qmi.order {
			qmi.function()
		}
	}

	if qm.AfterExecFunc != nil {
		qm.AfterExecFunc()
	}
}

func (qm QuickMenu) PromptLoop() {
	choice := 0

	for {
		if qm.BeforeMenuFunc != nil {
			qm.BeforeMenuFunc()
		}

		fmt.Printf("0. Quit\n")

		for _, qmi := range qm.items {
			fmt.Printf("%d. %s\n", qmi.order, qmi.description)
		}

		if qm.AfterMenuFunc != nil {
			qm.AfterMenuFunc()
		}

		var entry string
		n, err := fmt.Scanln(&entry)
		if n == 0 {
			// just refresh menu on <enter>
			continue
		}
		if err != nil {
			fmt.Printf("Wrong input: [%s]\n", err)
			continue
		}

		choice, err = strconv.Atoi(entry)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if choice > len(qm.items) || choice < 0 {
			fmt.Printf("Choice [%d] not available\n", choice)
			continue
		}

		if choice == 0 {
			break
		}

		if qm.BeforeExecFunc != nil {
			qm.BeforeExecFunc()
		}

		for _, qmi := range qm.items {
			if choice == qmi.order {
				qmi.function()
			}
		}

		if qm.AfterExecFunc != nil {
			qm.AfterExecFunc()
		}
	}
}
