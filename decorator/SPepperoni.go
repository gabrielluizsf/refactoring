package decorator

type PepperoniTopping struct{
  pizza IPizza
}

func (pt *PepperoniTopping) getPrice() int{
  pizzaPrice := pt.pizza.getPrice();
  return pizzaPrice + 5;
}
