package player_test

import (
	"liveboard/player"
	"testing"
)

func TestPlayerCreatesPlayerWithValidData(t *testing.T) {
	var pname player.PName = "John Smith"
	var phandicap player.PHandicap = 12
	validPlayer, err := player.Create([])

	want := "John Smith"
	got := validPlayer.Name()
	if err != nil {
		t.Error("player.Create should not throw an error.")
	}

	if want != got {
		t.Error("want %v != got %v", want, got)
	}

}
