package main

func rotLeft(a []int32, d int32) []int32 {
	aLength := int32(len(a))
	var shift int32

	if d >= aLength {
		shift = d % aLength
	} else {
		shift = d
	}

	return append(a[shift:len(a):len(a)], a[:shift]...)
}

func main() {}
