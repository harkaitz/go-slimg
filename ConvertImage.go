package slimg

import (
	"os/exec"
	"io"
	"bytes"
	"path/filepath"
)

// Image is a byte slice that holds the image data.
type Image []byte

// ConvertExecutable holds the name or location of the ImageMagick's
// convert binary.
var ConvertExecutable string = "convert"

func init() {
	// Check if ImageMagick is installed.
	var cmd    *exec.Cmd
	var stdout  bytes.Buffer
	var stderr  bytes.Buffer
	
	cmd = exec.Command(ConvertExecutable, "-version")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		panic("ImageMagick is not installed")
	}
	
	err = cmd.Wait()
	if err != nil {
		panic("ImageMagick is not installed")
	}
}

// ConvertImage converts an image to a new format and resizes it. It
// also limits the size of the output image.
func ConvertImage(dimXY, filename, field, outFormat string, outMaxMB int, fp io.Reader) (odata Image, err error) {
	var cmd    *exec.Cmd
	var stdout  bytes.Buffer
	var stderr  bytes.Buffer
	var suffix  string
	
	suffix = filepath.Ext(filename)
	switch suffix {
	case ".jpg",".jpeg",".png", ".JPG", ".JPEG", ".PNG":
	default: err = newErrorF(l("Unsupported file format, shall be jpg or png"), field); return
	}
	
	cmd = exec.Command(
		ConvertExecutable,
		"-resize", dimXY,
		"-strip", suffix[1:] + ":-",
		outFormat + ":-",
	)
	cmd.Stdin  = fp
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil { return }
	
	err = cmd.Wait()
	if err != nil { newErrorEF(l("Invalid image"), err, field); return }
	
	odata = stdout.Bytes()
	if len(odata) > outMaxMB * 1024 * 1024 {
		err = newErrorF(l("The image is too large"), field)
		return
	}
	
	return
}
