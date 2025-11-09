package puppyplace

import (
	gofarm "github.com/AboloreDev/GoFarm"
)
func Cock() string {
	return "kookoorookoo"
}

func Cocks() string {
	return "Loud kookoorookoo"
}

func BigBark() string {
	return gofarm.Sound(Cock())
}

func BigBarks() string {
	return gofarm.Sound(Cocks())
}
