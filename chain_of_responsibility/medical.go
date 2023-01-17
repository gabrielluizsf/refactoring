package chain_of_responsibility

type Medical struct{
  next Department
}

func (medi *Medical) execute(p *Patient){
  if p.medicineDone {
    printMessage("Medicine already given to Patient");
    medi.next.execute(p);
    return;
  }
  printMessage("Medical given medicine to Patient");
  p.medicineDone = true;
  medi.next.execute(p);
}

func (m *Medical) setNext(next Department){
  m.next = next;
}
