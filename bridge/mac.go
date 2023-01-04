package bridge

import "fmt"

type Mac struct {
    printer Printer
}

func (mac *Mac) Print() {
    fmt.Println("Print request for mac");
    mac.printer.PrintFile();
}

func (mac *Mac) SetPrinter(printer Printer) {
    mac.printer = printer;
}

