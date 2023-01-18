package command

func TVON_OR_TVOFF(){
  tv :=  &TV{};

  onCommand  := &OnCommand{
    device: tv,
  }
  offCommand := &OffCommand{
    device: tv,
  }
  onButton   := &Button{
    command: onCommand,
  }
  offButton  := &Button{
    command: offCommand,
  }

  onButton.press();
  offButton.press();
}
