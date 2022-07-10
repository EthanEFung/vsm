package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message string
	vsm     *vendingMachine
}

func newModel() *model {
	return &model{
		message: "",
		vsm:     newVendingMachine(1, 4),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var err error
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			m.message, err = m.vsm.addItem(1)
		case "r":
			m.message, err = m.vsm.requestItem()
		case "i":
			m.message, err = m.vsm.insertMoney(1)
		case "d":
			m.message, err = m.vsm.dispenseItem()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	if err != nil {
		m.message = err.Error()
	}
	return m, nil
}
func (m model) View() string {
	var s string
	s += "The Vending (state) Machine\n\n"
	s += fmt.Sprintln("items left: ", m.vsm.itemCount)
	s += fmt.Sprintln("price:", m.vsm.itemPrice)
	s += fmt.Sprintln("money inserted:", m.vsm.moneyCount)

	s += fmt.Sprintln("Issue the following commands:")
	s += fmt.Sprintln("(a)dd item - (r)equest item - (i)nsert money - (d)ispense item - (q)uit")
	s += fmt.Sprintf("\nstatus:\n%v", m.message)
	return s
}

func main() {
	program := tea.NewProgram(newModel())
	log.Fatal(program.Start())
}
