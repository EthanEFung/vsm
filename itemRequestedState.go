package main

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() (string, error) {
	return "", errExistingRequest
}

func (i *itemRequestedState) addItem(count int) (string, error) {
	return "", errDispenseInProgress
}

func (i *itemRequestedState) insertMoney(money int) (string, error) {
	i.vendingMachine.incrementMoneyCount(money)
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	if money < i.vendingMachine.itemPrice {
		return "", errInsufficientMoney
	}
	return "OK", nil
}

func (i *itemRequestedState) dispenseItem() (string, error) {
	return "", errNoMoneyInserted
}
