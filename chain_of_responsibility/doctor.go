package chain_of_responsibility

type Doctor struct{
  next Department
}

func (doctor *Doctor) execute(patient *Patient){
  if patient.doctorCheckUpDone {
    printMessage("Doctor checkup already done");
    doctor.next.execute(patient);
    return;
  }
  printMessage("Doctor checking Patient");
  patient.doctorCheckUpDone = true;
  doctor.next.execute(patient);
}

func (doctor *Doctor) setNext(next Department){
  doctor.next = next;
}
