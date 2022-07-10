package quickmenu

import (
	"fmt"
)

type QuickMenuItem struct {
	order       int
	description string
	function    func()
}

type QuickMenu []QuickMenuItem

func (qm *QuickMenu) Add(description string, f func()) {
	n := QuickMenuItem{len(*qm) + 1, description, f}
	*qm = append(*qm, n)
}

func (qm QuickMenu) Prompt() {
	choice := 0

	for _, qmi := range qm {
		fmt.Printf("%d. %s\n", qmi.order, qmi.description)
	}

	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Printf("Wrong input: [%s]\n", err)
	}

	for _, qmi := range qm {
		if choice == qmi.order {
			qmi.function()
		}
	}
}
