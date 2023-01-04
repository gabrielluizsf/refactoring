package factoryMethod;

func SelectGun(gunType string){
  gun, err := getGun(gunType);
    thisErrorEqualNil(err);
  printDetails(gun);
}
