package chain_of_responsibility

func HOSPITAL(){
  cashier := &Cashier{};
  
  medical := &Medical{};
  medical.setNext(cashier);

  doctor := &Doctor{};
  doctor.setNext(medical);

  receptionist := &Reception{};
  receptionist.setNext(doctor);

  patient := &Patient{name: "Gabriel"};
  //Patient visiting
  receptionist.execute(patient);
}
