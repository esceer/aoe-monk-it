package app

import (
	"fmt"

	"github.com/esceer/aoe-monk-it/internal/config"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func init() {
	robotgo.MouseSleep = 5
	robotgo.KeySleep = 5
}

func RegisterHooks(cfg *config.Config) {
	hook.Register(hook.KeyDown, cfg.Keystrokes.Record, func(e hook.Event) {
		fmt.Println("record mouse position...")
		recordPosition(cfg)
	})
	hook.Register(hook.KeyDown, cfg.Keystrokes.Deselect, func(e hook.Event) {
		fmt.Println("monk it!")
		deselectFromGroup(cfg)
	})
	fmt.Println("Listening to events...")
	s := hook.Start()
	<-hook.Process(s)
}

func deselectFromGroup(cfg *config.Config) {
	origX, origY := robotgo.GetMousePos()
	robotgo.Move(cfg.Coordinates.X, cfg.Coordinates.Y)
	robotgo.KeyDown("ctrl")
	robotgo.Click()
	robotgo.KeyUp("ctrl")
	robotgo.Move(origX, origY)
}

func recordPosition(cfg *config.Config) {
	cfg.Coordinates.X, cfg.Coordinates.Y = robotgo.GetMousePos()
	cfg.Save()
}
