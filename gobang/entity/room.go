package entity

type Room struct {
	Id         string        `json:"id"`
	Dialog     []DialogMsg   `json:"dialog"`
	Steps      []Chess       `json:"steps"`
	Started    bool          `json:"started"`
	Host       PlayerDetails `json:"host"`
	Challenger PlayerDetails `json:"challenger"`
	Spectators []Player      `json:"spectators"`
}
