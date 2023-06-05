package player

import "errors"

type PName string

const PnameEmpty string = "please provide a player name"
const PhandicapMaximum string = "maximum handicap is 54"

type PHandicap int8

type Player struct {
	name     PName
	handicap PHandicap
}

// function options may be better
func Create(...args) (Player, error) {
	newPlayer := Player{}
	for _, player := range args {
		switch v := player.(type) {
		case PName:
			err := NameValidate(v)
			if err != nil {
				return Player{}, err
			}
			newPlayer.name = v

		case PHandicap:
			err := HandicapValidate(v)
			if err != nil {
				return Player{}, err
			}
			newPlayer.handicap = v
		}

	}

	return newPlayer, nil
}

func NameValidate(name PName) error {
	if name == "" {
		return errors.New(PnameEmpty)
	}

	return nil
}
func HandicapValidate(handicap PHandicap) error {
	if handicap > 54 {
		return errors.New(PhandicapMaximum)
	}

	return nil
}
