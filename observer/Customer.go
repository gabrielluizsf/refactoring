package observer

import "fmt"

type Customer struct{
  id string
}

func (customer *Customer) update(itemName string){
  fmt.Printf("Sending email to customer %s for item %s\n", customer.id, itemName);

}

func (customer *Customer) getID() string{
  return customer.id;
}
