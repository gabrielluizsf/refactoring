package flyweight

import "github.com/theGOURL/warning"

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	playerDress, err := getDressFactorySingleInstance().getDressByType(dressType)
	warning.PRINT_DEFAULT_ERRORS(err, "ERROR: ")
	return &Player{
		playerType: playerType,
		dress:      playerDress,
	}
}
func (player *Player) newLocation(lat, long int) {
	player.lat = lat
	player.long = long
}
