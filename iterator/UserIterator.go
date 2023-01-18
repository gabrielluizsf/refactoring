package iterator

type UserIterator struct{
  index    int
  users []*User
}

func (ui *UserIterator) hasNext() bool{
  if ui.index < len(ui.users){
    return true
  }
  return false
}

func (ui *UserIterator) getNext() *User{
  if ui.hasNext(){
    user := ui.users[ui.index];
    ui.index++;
    return user;
  }
  return nil;
}
