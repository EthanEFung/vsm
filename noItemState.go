package main

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() (string, error) {
	return "", errOutOfStock
}

func (i *noItemState) addItem(count int) (string, error) {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return "item added", nil
}

func (i *noItemState) insertMoney(money int) (string, error) {
	return "", errOutOfStock
}

func (i *noItemState) dispenseItem() (string, error) {
	return "", errOutOfStock
}
