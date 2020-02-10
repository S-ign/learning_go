package gadget

import (
	"fmt"
)

//An interface allows you to interact with
//multiple types that you have created
//that have similar functionality by adding only
//the list of identical method names in the interface.
type Player interface {
	Play(string)
	Stop()
}


type TapePlayer struct {
	Batteries string
}
//////////////////////////////////////////////////
//SIMILAR METHODS
func (t TapePlayer) Play(song string) {
	fmt. Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt. Println("Stopped!")
}
//////////////////////////////////////////////////



type TapeRecorder struct {
	Microphones int
}
//////////////////////////////////////////////////
//SIMILAR METHODS
func (t TapeRecorder) Play(song string) {
	fmt. Println("Playing", song)
}
func (t TapeRecorder) Stop() {
	fmt. Println("Stopped!")
}
//////////////////////////////////////////////////
func (t TapeRecorder) Record() {
	fmt. Println("Recording")
}
