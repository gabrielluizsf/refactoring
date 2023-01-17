package chain_of_responsibility

import "fmt"

type Reception struct{
  next Department
}

func printMessage(message string){
  fmt.Println("\n",message);
}

func (receptionist *Reception) execute(patient *Patient){
  if patient.registrationDone {
    printMessage("Patient registration already done");
    receptionist.next.execute(patient);
    return;
  }
  printMessage("Reception registering Patient");
  patient.registrationDone = true;
  receptionist.next.execute(patient);
}

func (receptionist *Reception) setNext(next Department){
  receptionist.next = next;
}
