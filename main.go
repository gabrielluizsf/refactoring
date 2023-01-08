package main

import(
   "github.com/GabrielLuizSF/refactoring/abstractFactory"
   "github.com/GabrielLuizSF/refactoring/adapters"
   "github.com/GabrielLuizSF/refactoring/builder"
   "github.com/GabrielLuizSF/refactoring/bridge"
   "github.com/GabrielLuizSF/refactoring/factoryMethod"
   "github.com/GabrielLuizSF/refactoring/composite"
   "github.com/GabrielLuizSF/refactoring/prototype"
   "github.com/GabrielLuizSF/refactoring/decorator"
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
      designPattern(5, "Factory Method");
       factoryMethod.SelectGun("Lightsaber");
       factoryMethod.SelectGun("AK47");
      designPattern(6, "Composite");
       composite.Search();
      designPattern(7,"Prototype");
       prototype.CloneFileANDFolder();
      designPattern(8, "Decorator");
       decorator.SelectPizza();
        
}
