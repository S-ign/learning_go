package main

import (
	"learning_go/interfaces/gadget"
)

//Since we've created an interface between our similar types
//we can use this one function ( and one interface type )
//and use the similar methods of both of our types.
func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	player1 := gadget.TapePlayer{}
	playList(player1, mixtape)

	player2 := gadget.TapeRecorder{}
	playList(player2, mixtape)
}
