package image_keeper

import "testing"

func TestImageKeeper_CreateSource(t *testing.T) {
	i := New()
	i.CreateCompleteRepo("v0.0.2RC")
}

func TestImageKeeper_CreateCompleteImages(t *testing.T) {
	i := New()
	i.CreateCompleteImages("v0.0.2RC")
}
