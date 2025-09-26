package game

type NPCPlayer struct {
	*Player
	Trickster bool
}

func NewNPC(name string, isTrickster bool) *NPCPlayer {
	return &NPCPlayer{
		&Player{Name: name},
		isTrickster,
	}
}
