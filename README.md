# GO-SLIMG

This is a Go package for converting images to a smaller size and
embedding them in HTML pages. It is intended to be used in web
applications.

## Go documentation

    package slimg // import "github.com/harkaitz/go-slimg"
    
    var ConvertExecutable string = "convert"
    var DefaultLogoURL string = "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png"
    type Image []byte
        func ConvertImage(dimXY, filename, field, outFormat string, outMaxMB int, fp io.Reader) (odata Image, err error)
    type Logo []byte
        func ReadLogo(fp io.Reader, filename, field string, maxMB int) (b Logo, err error)
    type QR string
        func NewRandomQR() QR

## Go type QR

    package slimg // import "."
    
    type QR string
        QR is a string containing to be presented as a QR code. You can embed it in
        HTML pages throw templating using ImageSRC().
    
    func NewRandomQR() QR
    func (qr QR) ImageSRC() (qrCode template.URL)
    func (qr QR) String() string

## Go type Logo

    package slimg // import "."
    
    type Logo []byte
        Logo type holds a PNG file's content. You can embed this logo in HTML pages
        using ImageSRC() and templating.
    
    func ReadLogo(fp io.Reader, filename, field string, maxMB int) (b Logo, err error)
    func (l Logo) ImageSRC() (url template.URL)

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
