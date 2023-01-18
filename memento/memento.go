package memento

type Memento struct{
  state string
}

func (meme *Memento) getSavedState() string{
  return meme.state;
}
