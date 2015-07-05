/**
 * goss basic example
 */

package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/ww24/goss"
)

func main() {
	ss, err := goss.New()

	file, err := os.Create("screenshot.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to capture screenshot.")
		os.Exit(1)
	}
	defer file.Close()

	err = png.Encode(file, ss.ToImage())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to save screenshot.")
		os.Exit(1)
	}
}
