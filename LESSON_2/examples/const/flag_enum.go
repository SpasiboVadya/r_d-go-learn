package main

import (
	"fmt"
	_ "github.com/valyala/fasthttp"
	"lesson02/getter"

	"github.com/google/uuid"
)

const _ = getter.X

type FlagDayOfWeek uint8

const (
	FMonday FlagDayOfWeek = 1 << iota
	FTuesday
	FWednesday
	FThursday
	FFriday
	FSaturday
	FSunday
)

func (f FlagDayOfWeek) Contains(day FlagDayOfWeek) bool {
	return f&day != 0
}

func main() {
	fmt.Println(uuid.New())
}
