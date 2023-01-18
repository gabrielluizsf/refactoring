package iterator

type UserCollection struct{
  users []*User
}

func (uc *UserCollection) createIterator() Iterator{
  return &UserIterator{
    users: uc.users,
  }
}
