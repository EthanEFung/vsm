package main

import (
	"fmt"
)

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() (string, error) {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return "", errNoItemPresent
	}
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return "item requested", nil
}

func (i *hasItemState) addItem(count int) (string, error) {
	msg := fmt.Sprintf("%d items added\n", count)
	i.vendingMachine.incrementItemCount((count))
	return msg, nil
}

func (i *hasItemState) insertMoney(money int) (string, error) {
	return "", errNoItemSelected
}

func (i *hasItemState) dispenseItem() (string, error) {
	return "", errNoItemSelected
}
