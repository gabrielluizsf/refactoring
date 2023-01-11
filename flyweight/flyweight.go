package flyweight

import "fmt"

func StartGame(){
  game := newGame();

  game.addTerrorist(TerroristDressType);
  game.addTerrorist(TerroristDressType);
  game.addTerrorist(TerroristDressType);
  game.addTerrorist(TerroristDressType);
  game.addTerrorist(TerroristDressType);

  game.addHeroe(CounterTerroristDressType);
  game.addHeroe(CounterTerroristDressType);
  game.addHeroe(CounterTerroristDressType);
  game.addHeroe(CounterTerroristDressType);
  game.addHeroe(CounterTerroristDressType);

  dressFactoryInstance := getDressFactorySingleInstance();

  for dressType, dress := range dressFactoryInstance.dressMap {
    fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor());
  }
}
