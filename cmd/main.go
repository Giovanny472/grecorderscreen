package main

import (
	"bytes"
	"image/png"
	"log"

	"github.com/kbinani/screenshot"
	"golang.design/x/clipboard"
)

func main() {

	// screenshot
	bounds := screenshot.GetDisplayBounds(0)
	img, errScreen := screenshot.CaptureRect(bounds)
	if errScreen != nil {
		log.Fatalln("ошибка при screenshot: ", errScreen.Error())
	}

	// screen_image to bytes
	reqBodyBytes := new(bytes.Buffer)
	png.Encode(reqBodyBytes, img)

	// init clipboard
	errClip := clipboard.Init()
	if errClip != nil {
		log.Fatalln("ошибка при clipboard: ", errClip.Error())
	}

	// img to clipboard
	clipboard.Write(clipboard.FmtImage, reqBodyBytes.Bytes())
	clipboard.Read(clipboard.FmtImage)

}
