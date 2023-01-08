package singleton

import (
  "fmt"
  "sync"
)

type single struct{};

var (
  lock =  &sync.Mutex{}
  singleInstance *single
);

func getInstance() *single{
  if singleInstance != nil{
    fmt.Println("[created]");
  }else{
    lock.Lock();
    defer lock.Unlock();
    fmt.Println("Creating Single Instance  now");
    singleInstance = &single{}
  }
  return singleInstance;
}
