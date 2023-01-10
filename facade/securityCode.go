package facade

import (
  "fmt"
)

type SecurityCode struct{
  code int
}

func newSecurityCode(code int) *SecurityCode{
  return &SecurityCode{code: code};
}

func (sc *SecurityCode) checkCode(code int) error{
  if sc.code != code{
    fmt.Errorf("Access Denied");
  }
  fmt.Println("Allowed Access");
  return nil;
}
