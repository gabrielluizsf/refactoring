package decorator

type TomatoTopping struct{
  pizza IPizza
}

func (tt *TomatoTopping) getPrice() int{
  pizzaPrice := tt.pizza.getPrice();
  return pizzaPrice + 7;
}
