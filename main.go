package main

import(
   "github.com/GabrielLuizSF/refactoring/abstractFactory"
   "github.com/GabrielLuizSF/refactoring/adapters"
   "github.com/GabrielLuizSF/refactoring/builder"
   "github.com/GabrielLuizSF/refactoring/bridge"
)

func main(){
   print_message("Design Patterns && Refactoring");
     designPattern(1 ,"Abstract Factory");
      abstractFactory.Client();
     designPattern(2 ,"Adapters");
      adapters.Start();
     designPattern(3, "Builder");
      builder.MyBuilds();
     designPattern(4, "Bridge");
      bridge.PrintOut();
  

}
