package factoryMethod;

type LightSaber struct{
  Gun
}

func newLightSaber()IGun{
  return &LightSaber{
    Gun: Gun{
      name: "Lightsaber",
      power: 400,
    },
  }
}
