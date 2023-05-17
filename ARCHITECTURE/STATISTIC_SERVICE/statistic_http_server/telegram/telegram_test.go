package telegram

import "testing"

func TestDo(t *testing.T) {
	tg := New()
	tg.SendMessage("Hello")
}
