package bridge 

import "fmt"

type Windows struct {
    printer Printer
}

func (win *Windows) Print() {
    fmt.Println("Print request for windows");
    win.printer.PrintFile();
}

func (win *Windows) SetPrinter(printer Printer) {
    win.printer = printer;
}


