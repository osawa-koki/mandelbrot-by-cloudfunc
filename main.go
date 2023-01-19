package p

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	xmin, err := strconv.ParseFloat(r.FormValue("xmin"), 64)
	if err != nil {
		xmin = -2
	}

	ymin, err := strconv.ParseFloat(r.FormValue("ymin"), 64)
	if err != nil {
		ymin = -2
	}

	xmax, err := strconv.ParseFloat(r.FormValue("xmax"), 64)
	if err != nil {
		xmax = 2
	}

	ymax, err := strconv.ParseFloat(r.FormValue("ymax"), 64)
	if err != nil {
		ymax = 2
	}

	width, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		width = 1024
	}

	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		height = 1024
	}

	_iter, err := strconv.Atoi(r.FormValue("iter"))
	var iter uint8
	if err != nil {
		iter = 200
	} else {
		iter = uint8(_iter)
	}

	_thre, err := strconv.Atoi(r.FormValue("thre"))
	var thre uint8
	if err != nil {
		thre = 15
	} else {
		thre = uint8(_thre)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			color := (func() color.Color {
				var v complex128
				for n := uint8(0); n < iter; n++ {
					v = v*v + z
					if cmplx.Abs(v) > 2 {
						return color.Gray{255 - thre*n}
					}
				}
				return color.Black
			})()
			img.Set(px, py, color)
		}
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
