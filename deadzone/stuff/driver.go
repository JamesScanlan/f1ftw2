package main

import (
	"fmt"
)

type driver struct {
	name string
	team string
}

func newDriver(name string, team string) *driver {
	d := driver{name: name, team: team}
	return &d
}

func main() {
	d := driver{"Lewis Hamilton", "Ferrari"}

	fmt.Println(d.name)

	d2 := newDriver("Max Verstappen", "Red Bull")

	fmt.Println(d2.name)
}
