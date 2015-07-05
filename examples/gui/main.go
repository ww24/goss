/**
 * goss GUI example
 */

package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/themes/dark"
	"github.com/ww24/goss"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()

	// mx := source.Bounds().Max
	// fmt.Println(mx)

	window := theme.CreateWindow(500, 500, "Screenshot")
	size := window.Viewport().SizePixels()
	fmt.Println(size)
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	go func() {
		for {
			source, err := goss.Capture()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to capture screenshot.")
				os.Exit(1)
			}

			rgba := image.NewRGBA(source.Bounds())
			draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
			texture := driver.CreateTexture(rgba, 1)
			img.SetTexture(texture)
		}
	}()

	window.OnClose(driver.Terminate)
}

func main() {
	flag.Parse()
	gl.StartDriver(appMain)
}
