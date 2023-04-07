package abra

// #cgo CFLAGS: -I.
// #cgo CXXFLAGS: -I.
// #cgo LDFLAGS: -L. -ltestcpplib
// #include <plusik.h>
import "C"

func Plusik(a, b int) C.int {
	return C.plusik(1, 2)
}
