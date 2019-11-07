package hats

type Hat struct {
	color bool
}

func (hat Hat) toString() string {
	if hat.color == true {
		return "BLACK"
	} else {
		return "WHITE"
	}
}
