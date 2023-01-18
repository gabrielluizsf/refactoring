package command

type OffCommand struct {
  device Device
}

func (command *OffCommand) execute() {
  command.device.off();
}
