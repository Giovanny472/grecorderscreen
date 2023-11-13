package logic

import (
	"github.com/Giovanny472/grecordscreen/internal/model"
)

type screen struct {
	glf model.GlFw
}

var scr *screen

func NewScreen() model.Screen {
	if scr == nil {
		scr = &screen{glf: NewGlfw()}
	}
	return scr
}

func (s *screen) Show() {

	// инициализация
	s.glf.Init()

	// отображение окна для screenshot
	s.glf.ShowGLFW()

}
