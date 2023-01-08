package decorator

import "fmt"

func SelectPizza(){
  pizza := &KingMania{}

  pizzaWithPepperoni := &PepperoniTopping{
    pizza: pizza,
  }
  pizzaWithPepperoniANDTomato := &TomatoTopping{pizza: pizzaWithPepperoni}

  fmt.Println("Price of KingMania Pizza is: ",pizzaWithPepperoniANDTomato.getPrice());
}
