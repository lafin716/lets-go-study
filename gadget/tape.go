package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing : ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

