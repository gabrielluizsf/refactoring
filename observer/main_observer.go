package observer

func SendProductEmails(productName string){
  shirtItem := newItem(productName);

  observerFirst  := &Customer{id: "adajhasugfa@gmail.com"};
  observerSecond := &Customer{id: "xyz@asdf.com"};

  shirtItem.register(observerFirst);
  shirtItem.register(observerSecond);

  shirtItem.updateAvailability();
}
