package facade 

import "fmt"

type Account struct{
  name string
}

func newAccount(AccountName string) *Account{
  return &Account{name: AccountName};
}

func (acc *Account) checkAccount(AccountID string) error{
  if acc.name != AccountID{
    return fmt.Errorf("Access Denied");
  }
  fmt.Println("Account Verified");
  return nil;
}
