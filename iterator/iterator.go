package iterator

import "fmt"

func GetALLUsers(){
  CEO   := &User{
    name:    "CEO",
    age:      50,
    isAdmin:  true,
  }
  sysAdmin := &User{
    name:    "admin",
    age:      20,
    isAdmin:  true,
  }
  user  := &User{
    name:    "user",
    age:      18,
    isAdmin:  false,
  }
  userCollection := &UserCollection{
    users: []*User{CEO,sysAdmin,user},
  }
  iterator := userCollection.createIterator();

  for iterator.hasNext(){
    user := iterator.getNext();
    fmt.Printf("\nUser is %+v, is %+v years old, User is admin? %+v\n", user.name,user.age,user.isAdmin);
  }
}
