package flyweight

type game struct{
  terrorists []*Player
  heroes      []*Player
}

func newGame() *game{
  return &game{
    terrorists: make([]*Player, 1),
    heroes:     make([]*Player, 1),
  }
}

func (terror *game) addTerrorist(dressType string){
  player := newPlayer("T", dressType);
  terror.terrorists = append(terror.terrorists,player);
  return
}
func (counterTerrorist *game) addHeroe(dressType string){
  player := newPlayer("H", dressType);
  counterTerrorist.heroes = append(counterTerrorist.heroes,player);
  return
}

