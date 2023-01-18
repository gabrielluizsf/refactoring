package command

type OnCommand struct{
  device Device
}

func (command *OnCommand) execute() {
  command.device.on();
}
