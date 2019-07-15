package resolver

// #cgo CFLAGS: -I./softfloat/source/include
// #define SOFTFLOAT_FAST_INT64
// #include "softfloat.h"
//
// #cgo CXXFLAGS: -std=c++14
// #include "printqf.h"
// #include "print128.h"
import "C"
import (
	"fmt"
)

func PrintTest(){
	lo := uint64(1232)
	ho := uint64(0)
	ret := C.printi128(C.uint64_t(lo), C.uint64_t(ho))

	num := C.GoString(ret)
	fmt.Printf("envPrinti128 called result is:%v \r\n", num)
}