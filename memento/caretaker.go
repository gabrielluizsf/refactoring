package memento

type Caretaker struct{
  mementoArray []*Memento
}

func (caretaker *Caretaker) addMemento(meme *Memento){
  caretaker.mementoArray = append(caretaker.mementoArray, meme);
}

func (caretaker *Caretaker) getMemento(index int) *Memento{
  return caretaker.mementoArray[index];
}
