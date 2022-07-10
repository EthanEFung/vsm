package main

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount  int
	itemPrice  int
	moneyCount int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() (string, error) {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) (string, error) {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) (string, error) {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() (string, error) {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	v.itemCount += count
}

func (v *vendingMachine) incrementMoneyCount(money int) {
	v.moneyCount += money
}

func (v *vendingMachine) resetMoneyCount() {
	v.moneyCount = 0
}
