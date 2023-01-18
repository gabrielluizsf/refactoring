package mediator

func TrainStation(){
  stationManager := newStationManager();

  passenger   := &PassengerTrain{
    mediator: stationManager,
  }
  freightTrain := &FreightTrain{
    mediator: stationManager,
  }

  passenger.arrive();
  freightTrain.arrive();
  passenger.depart();
}
