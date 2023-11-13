package model

type GlFw interface {
	Init()
	ShowGLFW()
	ProcessingCoords(xs, ys, xe, ye float32)
	Clear()
	Close()
}
