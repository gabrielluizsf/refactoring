package factoryMethod;

import "fmt";

func thisErrorEqualNil(err error){
  if err != nil{
    fmt.Errorf("[ERROR]:",err);
  }
}


