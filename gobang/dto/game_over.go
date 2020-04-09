package dto

import "gobang/entity"

type GameOverDTO struct {
	RId    string               `json:"rid"`
	Winner entity.PlayerDetails `json:"winner"`
	Loser  entity.PlayerDetails `json:"loser"`
	Cause  string               `json:"cause"`
}
