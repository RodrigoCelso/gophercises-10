package game

import (
	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/player"
)

type Game struct {
	Shoe        deck.Deck
	DiscardTray deck.Deck
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
	shoeLast := len(g.Shoe) - 1
	if shoeLast < 5 {
		g.Shoe = append(g.Shoe, g.DiscardTray...)
		g.DiscardTray = deck.Deck{}
		g.Shoe = g.Shoe.Shuffle()
	}
	hitCard := g.Shoe[shoeLast]
	g.Shoe = g.Shoe[:shoeLast]
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
