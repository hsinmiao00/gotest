package main

import "fmt"
import "strconv"

func Decode(s string) int64 {
	// Check if input is valid
	if len(s)%8 != 0 {
		return -1
	}

	var result int64
	for i := len(s) - 8; i >= 0; i -= 8 {
		result <<= 7
		sb, _ := strconv.ParseInt(s[i:i+8], 2, 64)
		result |= (int64(sb) & int64(0x7F))
	}

	return result
}

func Encode(i int64) string {
	ret := ""
	for ; i > 0; i >>= 7 {
		t := 0x80 | (i & 0x7F)
		s := strconv.FormatInt(t, 2)
		ret += "0" + s[1:]
	}

	return "1" + ret[1:]
}

func main() {
	fmt.Println("300", Decode(Encode(300)))
	fmt.Println("3000", Decode(Encode(3000)))
	fmt.Println("30000", Decode(Encode(30000)))
}
