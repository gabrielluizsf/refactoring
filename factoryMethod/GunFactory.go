package factoryMethod;

import "fmt";

func getGun(gunType string)(IGun, error){
  if gunType =="AK47"{
   return newAK47(),nil;
  }
  if gunType == "Lightsaber"{
   return newLightSaber(), nil;
  }
  return nil, fmt.Errorf("Don't Have this gun type");
}
