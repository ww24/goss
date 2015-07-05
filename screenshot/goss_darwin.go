/**
 * screencapture package for OS X >= 10.8
 */

package screenshot

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import "goss_darwin.h"

void createScreenShot() {
  // screenshot
  CGImageRef screenShot = CGWindowListCreateImage(CGRectInfinite, kCGWindowListOptionOnScreenOnly, kCGNullWindowID, kCGWindowImageDefault);
  NSBitmapImageRep *bitmapRep = [[NSBitmapImageRep alloc] initWithCGImage:screenShot];
  CGImageRelease(screenShot);

  NSData *data = [bitmapRep representationUsingType:NSPNGFileType properties:nil];
  // save image
  // [data writeToFile:@"test.png" atomically:YES];

  // call golang function
  callback([data length], [data bytes]);
}
*/
import "C"
import "image"

var errC = make(chan error)
var imgC = make(chan image.Image)

// Capture method
func Capture() (img image.Image, err error) {
	go C.createScreenShot()

	err = <-errC
	if err != nil {
		return
	}

	img = <-imgC
	return
}
