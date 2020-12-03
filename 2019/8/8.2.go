package _2019

const (
	pixelColourBlack       = '0'
	pixelColourWhite       = '1'
	pixelColourTransparent = '2'
)

type layerWithPixels struct {
	pixels string
	width  int
}

func (p *layerWithPixels) getColorAt(i int, j int) rune {
	return rune(p.pixels[p.width*i+j])
}

type layers []*layerWithPixels

func (l layers) getColorAt(i int, j int) rune {
	for _, l := range l {
		c := l.getColorAt(i, j)
		if c != pixelColourTransparent {
			return c
		}
	}

	return pixelColourTransparent
}

type imageDecoder struct {
	layers        layers
	width, height int
}

func newImageDecoder(s string, width, height int) *imageDecoder {
	layerSize := width * height
	ls := make(layers, 0)
	for i := 0; i < len(s); i = i + layerSize {
		ls = append(ls, &layerWithPixels{pixels: s[i : i+layerSize], width: width})
	}
	return &imageDecoder{layers: ls, width: width, height: height}
}

func (d *imageDecoder) decode() string {
	result := ""
	for i := 0; i < d.height; i++ {
		for j := 0; j < d.width; j++ {
			c := d.layers.getColorAt(i, j)
			result += string(c)
		}
	}

	return result
}

func DecodeImage(s string, width, height int) string {
	d := newImageDecoder(s, width, height)

	return d.decode()
}
