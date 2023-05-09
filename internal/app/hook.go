package app

import (
	"fmt"

	"github.com/esceer/aoe-monk-it/internal/config"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	isLive = true
)

func init() {
	robotgo.MouseSleep = 10
	robotgo.KeySleep = 10
}

func RegisterHooks(cfg *config.Config) {
	hook.Register(hook.KeyDown, cfg.Keystrokes.Record, func(e hook.Event) {
		fmt.Println("record mouse position...")
		recordPosition(cfg)
	})
	hook.Register(hook.KeyDown, cfg.Keystrokes.Deselect, func(e hook.Event) {
		if isLive {
			fmt.Println("monk it!")
			deselectFromGroup(cfg)
		}
	})
	hook.Register(hook.KeyDown, cfg.Keystrokes.Toggle, func(e hook.Event) {
		switch isLive {
		case true:
			fmt.Println("toggle off")
		case false:
			fmt.Println("toggle on")
		}
		isLive = !isLive
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
