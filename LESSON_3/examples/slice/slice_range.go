package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	balance int
}

func (a Account) String() string {
	return strconv.Itoa(a.balance)
}

func AddBalance() {
	accounts := []Account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}
	for i := range accounts {
		accounts[i].balance += 1000
	}

	fmt.Printf("%+v", accounts)
}
