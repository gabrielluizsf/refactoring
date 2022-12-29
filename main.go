package main

import(
   "github.com/GabrielLuizSF/refactoring/abstractFactory"
   "github.com/GabrielLuizSF/refactoring/adapters"
)

func main(){
   print_message("Design Patterns && Refactoring");
     designPattern(1 ,"Abstract Factory");
      abstractFactory.Client();
     designPattern(2 ,"Adapters");
      adapters.Start();
  

}
