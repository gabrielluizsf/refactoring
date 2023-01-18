package memento

import "fmt"

func GetAllStatesAndRestoreStates(){
  caretaker  := &Caretaker{
    mementoArray: make([]*Memento, 0),
  }
  originator := &Originator{
    state: "INIT01",
  }
  printCurrentState(originator,caretaker);

  defineCurrentState(originator,"INIT02");
    printCurrentState(originator,caretaker);
  defineCurrentState(originator,"INIT03");
    printCurrentState(originator,caretaker);

  restoreToState(originator, caretaker, 1);
  restoreToState(originator, caretaker, 0);
}

func printCurrentState(origin *Originator,caretaker *Caretaker){
  fmt.Printf("Originator Current State: %s \n", origin.getState());
  caretaker.addMemento(origin.createMemento());
}

func defineCurrentState(origin *Originator, state string){
  origin.setState(state);
}

func restoreToState(origin *Originator,caretaker *Caretaker, address int){
  origin.restoreMemento(caretaker.getMemento(address));
  fmt.Printf("Restored to State: %s\n", origin.getState());
}
