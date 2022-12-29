package adapters

func Start() {

    client :=  &Client{};
    mac    :=  &Mac{};
    linux  :=  &Linux{};
      client.InsertLightningConnectorIntoComputer(mac);
      client.InsertLightningConnectorIntoComputer(linux);
    
    windowsMachine := &Windows{};
      windowsMachineAdapter := &WindowsAdapter{
          windowMachine: windowsMachine,
      };
      client.InsertLightningConnectorIntoComputer(windowsMachineAdapter);
}
