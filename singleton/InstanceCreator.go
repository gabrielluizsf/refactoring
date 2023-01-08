package singleton

import "fmt"

func CreateInstance(numberOfInstanceChecks int){
  for i := 0; i <= numberOfInstanceChecks; i ++{
    go getInstance()
  }
  fmt.Println("END");
}
