package entity

//在某个房间里的玩家详细信息
type PlayerDetails struct {
	Player
	Role  string `json:"role"`
	Color int8   `json:"color"`
	Ready bool   `json:"ready"`
}
