package qrcode

import (
	"github.com/skip2/go-qrcode"
	"image/color"
)

func NewQRCode() error {
	err := qrcode.WriteColorFile("https://www.baidu.com", qrcode.Medium, 256, color.Black, color.White, "baidu.png")
	if err != nil {
		return err
	}
	return nil
}