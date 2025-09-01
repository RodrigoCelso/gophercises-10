package game

import (
	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/player"
)

type Game struct {
	Shoe        []deck.Card
	DiscardTray []deck.Card
	Dealer      *player.Player
	Players     []*player.Player
}

func (g *Game) StartRound() {
	for _, player := range g.Players {
		g.Hit(player)
		g.Hit(player)
	}
	g.Hit(g.Dealer)
	g.Hit(g.Dealer)
}

func (g *Game) Hit(p *player.Player) deck.Card {
	shoeLen := len(g.Shoe)
	hitCard := g.Shoe[shoeLen-1]
	g.Shoe = g.Shoe[:shoeLen-1]
	p.Hand = append(p.Hand, hitCard)
	return hitCard
}

func (g *Game) Stand(p *player.Player) {

}

func (g *Game) DoubleDown(p *player.Player) {

}

func (g *Game) Split(p *player.Player) {

}

func (g *Game) ClearDiscardTray(p *player.Player) {

}
