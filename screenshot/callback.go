/**
 * callback method
 */

package screenshot

// #include <stdlib.h>
import "C"
import (
	"image/png"
	"math"
	"reflect"
	"unsafe"
)

type reader struct {
	cur  int
	data []byte
}

// Read method
func (r *reader) Read(p []byte) (n int, err error) {
	if r.cur == len(r.data)-1 {
		return
	}

	copy(p, r.data[r.cur:])
	n = int(math.Min(float64(len(r.data)-r.cur), float64(len(p))))
	r.cur += n

	return
}

//export callback
func callback(length uint, bytes unsafe.Pointer) {
	r := &reader{}
	sheader := (*reflect.SliceHeader)(unsafe.Pointer(&r.data))
	sheader.Data = uintptr(bytes)
	sheader.Len = int(length)

	img, err := png.Decode(r)
	if err != nil {
		errC <- err
		return
	}

	errC <- nil
	imgC <- img
}
