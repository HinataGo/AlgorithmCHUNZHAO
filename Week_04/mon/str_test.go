package mon

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	str1, str2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	res := strStr(str1, str2)
	fmt.Println(res)
}
