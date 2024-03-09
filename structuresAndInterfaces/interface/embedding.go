package _interface

import "fmt"

type Printers interface {
	Print()
}

type Scanners interface {
	Scan()
}

// This means any type that satisfies the
// PrinterAndScanner interface must implement
// both the Print() and Scan() methods.

type PrintersAndScanner interface {
	Printers
	Scanners
}

type PrinterScannerImpl struct{}

func (ps PrinterScannerImpl) Print() {
	fmt.Println("Printing")
}

func (ps PrinterScannerImpl) Scan() {
	fmt.Println("Scanning")
}

func PrintAndScan(ps PrintersAndScanner) {
	ps.Print()
	ps.Scan()
}
