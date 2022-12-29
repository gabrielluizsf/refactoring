package adapters

type Linux struct{}

func (tux *Linux) InsertIntoLightningPort(){
  printDefaultMessage("Linux");
}
