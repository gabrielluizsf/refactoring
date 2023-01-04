package factoryMethod;

import "fmt";

func printDetails(gun IGun){
  fmt.Printf("Gun: %s\n", gun.getName());
  fmt.Printf("Power: %d\n\n", gun.getPower());
}
