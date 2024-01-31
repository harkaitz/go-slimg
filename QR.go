package slimg

import (
	"log"
	"github.com/skip2/go-qrcode"
	"encoding/base64"
	"html/template"
	"github.com/google/uuid"
)

// QR is a string containing to be presented as a QR code. You can
// embed it in HTML pages throw templating using ImageSRC().
type QR string

// NewRandomQR creates a QR containing a random UUID.
func NewRandomQR() (QR) {
	var err error
	var u   uuid.UUID
	u, err = uuid.NewRandom()
	if err != nil { log.Panic(err) }
	return QR(u.String())
}

// ImageSRC returns a "data:image/png" to be putted in (html img src="").
func (qr QR) ImageSRC() (qrCode template.URL) {
	
	var png   []byte
	var err     error
	
	png, err = qrcode.Encode(string(qr), qrcode.Medium, 256)
	if err != nil { log.Panic(err); return "" }
	
	return template.URL(
		"data:image/png;charset=utf-8;base64," +
		base64.StdEncoding.EncodeToString(png),
	)
}

// String returns the string contained in the QR.
func (qr QR) String() string {
	return string(qr)
}
