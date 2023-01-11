package flyweight

type CounterTerroristDress struct{
  color string
}

func (CT *CounterTerroristDress) getColor() string{
  return CT.color;
}

func newCounterTerroristDress() *CounterTerroristDress{
  return &CounterTerroristDress{color: "green"};
}
