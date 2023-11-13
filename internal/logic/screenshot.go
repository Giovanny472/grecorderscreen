package logic

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/Giovanny472/grecordscreen/internal/model"
	"github.com/kbinani/screenshot"
	"golang.design/x/clipboard"
)

type scrshot struct {
	rect   image.Rectangle
	buffSc []byte
}

var scsh *scrshot

func NewScreenshot() model.Screenshot {
	if scsh == nil {
		scsh = new(scrshot)
	}
	return scsh
}

func (s *scrshot) SetRectScreen(xstart, ystart, xend, yend int) {
	s.rect = image.Rectangle{Min: image.Point{X: xstart, Y: ystart}, Max: image.Point{X: xend, Y: yend}}
}

func (s *scrshot) DoScreenshot() {

	img, err := screenshot.CaptureRect(s.rect)
	if err != nil {
		log.Fatal("невозможно сделать screenshot: ", err, ", Rect: ", s.rect)
	}

	bufImg := new(bytes.Buffer)
	err = png.Encode(bufImg, img)
	if err != nil {
		log.Fatal("невозможно encode screenshot:", err)
	}

	s.buffSc = bufImg.Bytes()
}

func (s *scrshot) ScreenshotToClipboard() {
	clipboard.Write(clipboard.FmtImage, s.buffSc)
}
