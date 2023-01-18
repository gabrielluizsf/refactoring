package command

type Button struct{
  command Command
}

func (button *Button) press(){
  button.command.execute();
}
