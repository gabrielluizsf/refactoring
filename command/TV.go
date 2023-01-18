package command

import "fmt"

type TV struct{
  isRunning bool
}

func (tv *TV) on() {
  tv.isRunning = true;
  fmt.Println("\nTurning TV on");
}

func (tv *TV) off() {
  tv.isRunning = false;
  fmt.Println("\nTurning TV off");
}
