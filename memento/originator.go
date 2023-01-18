package memento

type Originator struct{
  state string
}

func (origin *Originator) createMemento() *Memento{
  return &Memento{state: origin.state};
}

func (origin *Originator) restoreMemento(meme *Memento){
  origin.state = meme.getSavedState();
}

func (origin *Originator) setState(state string){
  origin.state = state;
}

func (origin *Originator) getState() string{
  return origin.state;
}
