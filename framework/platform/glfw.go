package platform

import (
	"image"
	"strconv"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/wieku/danser-go/framework/assets"
	"github.com/wieku/danser-go/framework/graphics/texture"
)

var iconSizes = []int{128, 48, 24, 16}

func LoadIcons(win *glfw.Window, prefix, suffix string) {
	var pixmaps []*texture.Pixmap
	var images []image.Image

	for _, size := range iconSizes {
		pxMap, err := assets.GetPixmap("assets/textures/" + prefix + strconv.Itoa(size) + suffix + ".png")
		if err != nil {
			continue
		}

		pixmaps = append(pixmaps, pxMap)
		images = append(images, pxMap.NRGBA())
	}

	win.SetIcon(images)

	for _, pxMap := range pixmaps {
		pxMap.Dispose()
	}
}
