package chain_of_responsibility

type Cashier struct{
  next Department
}

func (c *Cashier) execute(p *Patient){
  if p.paymentDone {
    printMessage("Payment Done");
  }
  printMessage("Cashier Getting Money From Patient Patient");
}

func (c *Cashier) setNext(next Department){
  c.next = next;
}
