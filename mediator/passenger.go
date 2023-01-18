package mediator

import "fmt"

type PassengerTrain struct{
  mediator Mediator
}

func (people *PassengerTrain) arrive(){
  if !people.mediator.canArrive(people){
    fmt.Println("Passenger Train: Arrival blocked, please waiting");
    return
  }
  fmt.Println("Passenger Train: Arrived");
}

func (people *PassengerTrain) depart(){
  fmt.Println("Passenger Train: Leaving");
  people.mediator.notifyAboutDeparture();
}

func (people *PassengerTrain) permitArrival() {
  fmt.Println("Passenger Train: Arrival permitted, arriving");
  people.arrive();
}
