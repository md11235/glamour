package ansi

import (
	"github.com/mattn/go-sixel"
	"image"
	"io"
	"os"
)

import _ "image/jpeg"
import _ "image/png"

// An ImageElement is used to render images elements.
type ImageElement struct {
	Text    string
	BaseURL string
	URL     string
	Child   ElementRenderer // FIXME
}

func (e *ImageElement) Render(w io.Writer, ctx RenderContext) error {
	if len(e.Text) > 0 {
		el := &BaseElement{
			Token: e.Text,
			Style: ctx.options.Styles.ImageText,
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}
	if len(e.URL) > 0 {
		f, err := os.Open(e.URL)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// Calling the generic image.Decode() will tell give us the data
		// and type of image it is as a string. We expect "png"
		imageData, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		//fmt.Println(imageData)
		// fmt.Println(imageType)

		// sixel.NewEncoder(os.Stdout).Encode(imageData)

		sixel.NewEncoder(w).Encode(imageData)

		//el := &BaseElement{
		//	Token:  resolveRelativeURL(e.BaseURL, e.URL),
		//	Prefix: " ",
		//	Style:  ctx.options.Styles.Image,
		//}
		//err := el.Render(w, ctx)
		//if err != nil {
		//	return err
		//}

	}

	return nil
}
