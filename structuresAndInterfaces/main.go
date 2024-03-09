package main

import (
	"base/structuresAndInterfaces/interface"
	"fmt"
)

type Printer interface {
	Print()
}

type Person struct {
	Name string
}

func (p Person) Print() {
	fmt.Println(p.Name)
}

func PrintPerson(p Printer) {
	p.Print()
}

// Person satisfies the Printer interface
// that is why we passed in parameter Printer
// type argument

func main() {
	//p := Person{Name: "John doe"}
	//PrintPerson(p)
	//
	//ps := _interface.PrinterScannerImpl{}
	//_interface.PrintAndScan(ps)

	cal := _interface.SimpleCalculator{}
	_interface.PerformCalculation(10, 2, cal)

	_interface.PrintAny(42)
	_interface.PrintAny("Hello")
	_interface.PrintAny(true)
}
