package logic

import (
	"github.com/Giovanny472/grecordscreen/internal/model"
)

type app struct {
	screen model.Screen
}

var ap *app

func NewApp() model.IApp {
	if ap == nil {
		ap = &app{screen: NewScreen()}
	}
	return ap
}

func (a *app) Exec() {
	a.screen.Show()
}
