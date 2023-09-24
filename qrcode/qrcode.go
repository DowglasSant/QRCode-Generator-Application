package qrcode

import (
	"bytes"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func createQRCode(dataString string) ([]byte, error) {
	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		return nil, err
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, qrCode)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func SaveQRCodeImage(dataString string) ([]byte, error) {
	fileBytes, err := createQRCode(dataString)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("qrCode.png", fileBytes, 0644)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}
