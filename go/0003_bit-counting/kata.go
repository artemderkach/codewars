// https://www.codewars.com/kata/bit-counting/train/go
package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func CountBits(n uint) int {
	b := strconv.FormatUint(uint64(n), 2)
	count := strings.Count(fmt.Sprintf(b), "1")
	return count
}
