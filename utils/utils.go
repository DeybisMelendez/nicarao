package utils

import "math/bits"

func BitscanForward(x uint64) int {
	return bits.TrailingZeros64(x)
}

func BitscanReverse(x uint64) int {
	return 63 - bits.LeadingZeros64(x)
}
