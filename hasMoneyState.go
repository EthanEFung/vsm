package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() (string, error) {
	return "", errDispenseInProgress
}

func (i *hasMoneyState) addItem(count int) (string, error) {
	return "", errDispenseInProgress
}

func (i *hasMoneyState) insertMoney(money int) (string, error) {
	i.vendingMachine.incrementMoneyCount(money)
	if i.vendingMachine.moneyCount < i.vendingMachine.itemPrice {
		return "", errInsufficientMoney
	}
	return "OK", nil
}

func (i *hasMoneyState) dispenseItem() (string, error) {
	msg := fmt.Sprintln("item dispensed")
	if i.vendingMachine.moneyCount < i.vendingMachine.itemPrice {
		return "", errInsufficientMoney
	}

	if i.vendingMachine.moneyCount > i.vendingMachine.itemPrice {
		msg += fmt.Sprintln("here is your change: ", i.vendingMachine.moneyCount-i.vendingMachine.itemPrice)
	}

	i.vendingMachine.resetMoneyCount()

	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1

	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return msg, nil
	}
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return msg, nil
}
