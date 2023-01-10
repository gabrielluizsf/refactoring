package facade 

import "fmt"

type Notification struct{};

func (notify *Notification) sendWalletCreditNotification(){
  fmt.Println("Sending Wallet Credit Notification");
}

func (notify *Notification) sendWalletDebitNotification(){
  fmt.Println("Sending Wallet Debit Notification");
}
