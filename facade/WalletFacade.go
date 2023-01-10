package facade

import "fmt"

type WalletFacade struct{
  account      *Account
  wallet       *Wallet
  securityCode *SecurityCode
  notification *Notification
  ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade{
  fmt.Println("Starting create account");
  walletFacade := &WalletFacade{
    account:      newAccount(accountID),
    securityCode: newSecurityCode(code),
    wallet:       newWallet(),
    notification: &Notification{},
    ledger:       &Ledger{},
  }
  fmt.Println("[Account created] AccountID: ",accountID);
  return walletFacade;
}

func (wf *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error{
  fmt.Println("Starting add money to wallet");
  check := wf.account.checkAccount(accountID);
    returnError(check);
  checksecurityCode := wf.securityCode.checkCode(securityCode);
    returnError(checksecurityCode);
  wf.wallet.creditBalance(amount);
  wf.notification.sendWalletCreditNotification();
  wf.ledger.makeEntry(accountID,"credit",amount);
  return nil
}

func returnError(err error) error{
  if err != nil{
    return err;
  }
  return nil;
}

func (wf *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error{
  fmt.Println("Starting debit money from wallet");
  check := wf.account.checkAccount(accountID);
    returnError(check);
  checksecurityCode := wf.securityCode.checkCode(securityCode);
    returnError(checksecurityCode);
  debit := wf.wallet.debitBalance(amount);
    returnError(debit);
  wf.notification.sendWalletDebitNotification();
  wf.ledger.makeEntry(accountID,"credit",amount);
  return nil;
}
