package player

type PlayerName string
type PlayerHandicap int8

type Player struct {
	name     PlayerName
	handicap PlayerHandicap
}
