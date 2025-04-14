package main

import "fmt"

type Account struct {
	balance int
}

var hashSet map[int]struct{}

func AddBalance() {
	accounts := map[int]Account{
		1: {balance: 100},
		2: {balance: 200},
		3: {balance: 300},
	}

	v, ok := accounts[1]
	if ok {
		fmt.Printf("balance %+v", v)
	} else {
		fmt.Println("balance is nil")
	}

	//for k, v := range accounts {
	//	v.balance += 1000
	//	accounts[k] = v
	//}
	//
	//fmt.Printf("%+v", accounts)
}
