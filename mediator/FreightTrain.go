package mediator

import "fmt"

type FreightTrain struct{
  mediator Mediator
}

func (people *FreightTrain) arrive(){
  if !people.mediator.canArrive(people){
    fmt.Println("Freight Train: Arrival blocked, please waiting");
    return;
  }
  fmt.Println("Freight Train: Arrived");
}

func (people *FreightTrain) depart(){
  fmt.Println("Freight Train: Leaving");
  people.mediator.notifyAboutDeparture();
}

func (people *FreightTrain) permitArrival(){
  fmt.Println("Freight Train: Arrival permitted");
  people.arrive();
}
