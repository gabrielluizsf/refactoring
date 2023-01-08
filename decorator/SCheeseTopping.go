package decorator

type CheeseTopping struct{
  pizza IPizza
}

func (ct *CheeseTopping) getPrice() int{
  pizzaPrice := ct.pizza.getPrice();
  return pizzaPrice + 10;
}
