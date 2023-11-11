package model

type Screenshot interface {
	SetRectScreen(xstart, ystart, xend, yend int)
	DoScreenshot()
	ScreenshotToClipboard()
}
