package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var bitNum float64
var key string = "0x"

var tempkey string = "0x"
var isCracked = false
var wrongKey []float64

var a int = 0
var b int = 15

func main() {

	fmt.Println("Enter size of key in bites: ")
	fmt.Scanf("%f", &bitNum)
	keys := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bitNum)), nil)
	fmt.Println(keys)
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	for n := 0; float64(n) < bitNum; n++ {

		i := a + rand.Intn(b-a+1) // a ≤ n ≤ b

		if i == 0 {
			key += "0"
		} else if i == 1 {
			key += "1"
		} else if i == 2 {
			key += "2"
		} else if i == 3 {
			key += "3"
		} else if i == 4 {
			key += "4"
		} else if i == 5 {
			key += "5"
		} else if i == 6 {
			key += "6"
		} else if i == 7 {
			key += "7"
		} else if i == 8 {
			key += "8"
		} else if i == 9 {
			key += "9"
		} else if i == 10 {
			key += "A"
		} else if i == 11 {
			key += "B"
		} else if i == 12 {
			key += "C"
		} else if i == 13 {
			key += "D"
		} else if i == 14 {
			key += "E"
		} else if i == 15 {
			key += "F"
		}

	}
	fmt.Println(key)

	duration := time.Since(start)
	fmt.Println(duration)
}
