package slimg

import (
	"io"
	"encoding/base64"
	"html/template"
)

// Logo type holds a PNG file's content. You can embed this logo
// in HTML pages using ImageSRC() and templating.
type Logo []byte

// DefaultLogoURL is the default logo image to show if a Logo is
// empty or uninitialized.
var DefaultLogoURL string = "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png"

// ReadLogo reads a file and converts it to "256x" "png" format. It
// uses ImageMagick's convert command for that.
func ReadLogo(fp io.Reader, filename, field string, maxMB int) (b Logo, err error) {
	var bb []byte
	bb, err = ConvertImage("256x", filename, field, "png", 2, fp)
	return Logo(bb), err
}

// ImageSRC returns a "data:image/png" to be putted in (html img src="").
func (l Logo) ImageSRC() (url template.URL) {
	if len(l) == 0 {
		return template.URL(DefaultLogoURL)
	}
	return template.URL(
		"data:image/png;base64," +
		base64.StdEncoding.EncodeToString(l),
	)
}
