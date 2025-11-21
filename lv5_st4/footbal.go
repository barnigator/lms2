package main

import (
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		name,
		goals,
		misses,
		assists,
		0,
	}

	p.calculateRating()

	return p
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		p1 := players[i].Goals
		p2 := players[j].Goals
		if p1 != p2 {
			return p1 > p2
		} else {
			return players[i].Name < players[j].Name
		}
	})

	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		p1 := players[i].Rating
		p2 := players[j].Rating
		if p1 != p2 {
			return p1 > p2
		} else {
			return players[i].Name < players[j].Name
		}
	})

	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		p1 := players[i].Goals
		p2 := players[j].Goals
		if players[i].Misses != 0 {
			p1 = players[i].Goals / players[i].Misses
		}
		if players[j].Misses != 0 {
			p2 = players[j].Goals / players[j].Misses
		}
		if p1 != p2 {
			return p1 > p2
		} else {
			return players[i].Name < players[j].Name
		}
	})
	return players
}
