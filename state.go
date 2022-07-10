package main

import "errors"

var (
	errNoItemPresent      = errors.New("no item present")
	errNoItemSelected     = errors.New("please select an item")
	errExistingRequest    = errors.New("item already requested")
	errDispenseInProgress = errors.New("item dispense in progress")
	errInsufficientMoney  = errors.New("requires more money")
	errNoMoneyInserted    = errors.New("please insert money")
	errOutOfStock         = errors.New("item out of stock")
)

type state interface {
	addItem(int) (string, error)
	requestItem() (string, error)
	insertMoney(money int) (string, error)
	dispenseItem() (string, error)
}
