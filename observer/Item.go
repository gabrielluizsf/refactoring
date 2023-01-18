package observer

import "fmt"

type Item struct{
  observerList []Observer
  name           string
  inStock        bool
}

func newItem(name string) *Item{
  return &Item{
    name: name,
  }
}

func (item *Item) updateAvailability(){
  fmt.Printf("Item %s is now in Stock\n",item.name);
  item.inStock = true;
  item.notifyAll();
}

func (item *Item) register(observer Observer){
  item.observerList = append(item.observerList, observer);
}

func (item *Item) deregister(observer Observer){
  item.observerList = removeFromSlice(item.observerList, observer);
}

func (item *Item) notifyAll(){
  for _, observer := range item.observerList{
    observer.update(item.name);
  }
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer{
  observerListLength := len(observerList);

  for indice, observer := range observerList{
    if observerToRemove.getID() == observer.getID(){
      observerList[observerListLength - 1],observerList[indice] = observerList[indice],observerList[observerListLength -1];
      return observerList[:observerListLength-1]
    }
  }
  return observerList;
}
