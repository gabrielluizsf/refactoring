package flyweight

import "fmt"

type DressFactory struct{
  dressMap  map[string]Dress
}

var (
  dressFactorySingleInstance = &DressFactory{
    dressMap: make(map[string]Dress),
  }
)

const (
  TerroristDressType        = "TDRESS"
  CounterTerroristDressType = "CTDRESS"
)

func (df *DressFactory) getDressByType(dressType string) (Dress,error){
  if df.dressMap[dressType] != nil{
    return df.dressMap[dressType], nil;
  }
  if dressType == TerroristDressType {
    df.dressMap[dressType] = newTerroristDress();
    return df.dressMap[dressType], nil;
  }
  if dressType == CounterTerroristDressType {
    df.dressMap[dressType] = newCounterTerroristDress();
    return df.dressMap[dressType], nil;
  }
  return nil, fmt.Errorf("Wrong Dress Type Passed");
}

func getDressFactorySingleInstance() *DressFactory {
  return dressFactorySingleInstance;
}
