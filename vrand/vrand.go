package goutils

import (
	"time"
	"math/rand"
)

func vGetRand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r
}

// return a int type rand num
func VrandInt(maxNum int) int {
	return vGetRand().Intn(maxNum)
}

// return a int32 type rand num
func VrandInt32(maxNum int32) int32 {
	return vGetRand().Int31n(maxNum)
}

// return a int64 type rand num
func VrandInt64(maxNum int64) int64 {
	return vGetRand().Int63n(maxNum)
}

// return a int slice max element is maxElement
func VrandIntSlice(maxElement int) []int {
	return vGetRand().Perm(maxElement)
}

//shuffle a slice
func Vshuffle() {
	
}

