package bridge 

import "fmt"

type Epson struct {
}

func (ep *Epson) PrintFile() {
    fmt.Println("Printing by a EPSON Printer");
}
