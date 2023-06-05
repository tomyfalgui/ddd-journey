package player_test

import (
	"liveboard/player"
	"testing"
)

func TestPlayerCreatesPlayerWithValidData(t *testing.T) {
	validPlayer, err := player.Create(player.CreateName("John Smith"))

	want := "John Smith"
	got := validPlayer.Name()
	if err != nil {
		t.Error("player.Create should not throw an error.")
	}

	if want != got {
		t.Error("want %v != got %v", want, got)
	}

}
