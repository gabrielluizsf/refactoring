package flyweight

type TerroristDress struct{
  color string
}

func (TD  *TerroristDress) getColor() string{
  return TD.color;
}

func newTerroristDress() *TerroristDress{
  red := "red"
  return &TerroristDress{color: red};
}
