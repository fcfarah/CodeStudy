package main

import (
	"fmt"
	"math"
)

func main() {
	//Value range of different int types
	fmt.Println("value range of different int types:")
	//fmt. Println ("int:", math. MININT, "~", math. Maxint) reports an error without math MinInt math. MaxInt
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	fmt.Println()
}
