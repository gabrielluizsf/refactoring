package facade

import "fmt"

type Ledger struct{};

func (ldg *Ledger) makeEntry(accountID, txnType string, amount int){
  fmt.Printf("Make ledger entry for account ID  %s with txnType %s for amount %d\n",accountID,txnType,amount);
  return
}
