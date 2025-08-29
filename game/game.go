package game

import (
	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/player"
)

type GameActions interface {
	Hit(*player.Player)
	Stand(*player.Player)
	DoubleDown(*player.Player)
	Split(*player.Player)
	ClearDeck(*player.Player)
}

type Game struct {
	Shoe        []deck.Card
	DiscardTray []deck.Card
	Players     []*player.Player
}

func (g *Game) Hit(p *player.Player) {
	shoeLen := len(g.Shoe)
	hitCard := g.Shoe[shoeLen-1:]
	g.Shoe = g.Shoe[:shoeLen-1]
	p.Deck = append(p.Deck, hitCard...)
}

func (g *Game) Stand(p *player.Player) {

}

func (g *Game) DoubleDown(p *player.Player) {

}

func (g *Game) Split(p *player.Player) {

}

func (g *Game) ClearDeck(p *player.Player) {

}
