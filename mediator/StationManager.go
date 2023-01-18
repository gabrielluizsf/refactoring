package mediator

type StationManager struct{
  isPlataformFree bool
  trainQueue    []Train
}

func newStationManager() *StationManager{
  return &StationManager{
    isPlataformFree: true,
  }
}

func (sm *StationManager) canArrive(t Train) bool{
  if !sm.isPlataformFree{
    sm.isPlataformFree = false;
    return true;
  }
  sm.trainQueue = append(sm.trainQueue, t);
  return false;
}

func (sm *StationManager) notifyAboutDeparture(){
  if !sm.isPlataformFree{
    sm.isPlataformFree = true;
  }
  if len(sm.trainQueue) > 0{
    firstTrainInQueue := sm.trainQueue[0];
    sm.trainQueue      = sm.trainQueue[1:];
    firstTrainInQueue.permitArrival();
  }
}
