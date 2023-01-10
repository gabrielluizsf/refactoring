package facade

import (
  "log"
)

func MyWallet(accountID string, securityCode int){
  walletFacade := newWalletFacade(accountID,securityCode);

  add := walletFacade.addMoneyToWallet(accountID,securityCode,5000);
    walletDefaultError(add);
  remove := walletFacade.deductMoneyFromWallet(accountID,securityCode,700);
    walletDefaultError(remove);
  remove7e3 := walletFacade.deductMoneyFromWallet(accountID,securityCode,7000);     
    walletDefaultError(remove7e3);
}

func walletDefaultError(err error){
  if err != nil{
    log.Fatalf("ERROR: %s\n",err.Error());
  }
}
